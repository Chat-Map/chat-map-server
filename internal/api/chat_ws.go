package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Chat-Map/chat-map-server/internal/application"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type room struct {
	id        int64
	clients   map[*websocket.Conn]bool
	broadcast chan message
}

type message struct {
	ChatID  int64  `json:"chat_id"`
	UserID  int64  `json:"sender_id"`
	Content string `json:"content"`
}

var chatRooms = make(map[int64]*room)
var roomMx = make(map[int64]*sync.Mutex)

func roomLock(roomID int64) {
	if roomMx[roomID] == nil {
		roomMx[roomID] = &sync.Mutex{}
	}
	roomMx[roomID].Lock()
}

func roomUnlock(roomID int64) {
	roomMx[roomID].Unlock()
}

func (s *Server) storeMessage(chatID int64, message message) (int64, error) {
	return s.uc.MessageStore.Execute(context.TODO(), application.StoreMessageCommandRequest{
		ChatID:  chatID,
		UserID:  message.UserID,
		Content: message.Content,
	})
}

// ChatWS godoc
//
//	@Summary		Connect to chat room with websockets
//	@Description	Conncects the user to the chat room to get notified of new messages
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			params				path		int64	true	"Chat ID"
//	@Success		200					{object}	message
//	@Failure		400,401,403,404,500	{object}	api.Response{data=interface{}}
//	@Security		bearerAuth
//	@Router			/chat/ws/{id} [get]
func (s *Server) chatws(w http.ResponseWriter, r *http.Request) {
	// Get chatID from var
	vars := mux.Vars(r)
	id := vars["id"]
	chatID, err := strconv.Atoi(id)
	if err != nil || chatID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid chat id or not provided"))
		return
	}
	roomID := int64(chatID)
	// TODO: check user is a chat member before open websocket connection

	// Load room (Or create if doesn't exist)
	roomLock(roomID)
	chatRoom, ok := chatRooms[roomID]
	if !ok {
		chatRoom = &room{
			id:        roomID,
			clients:   make(map[*websocket.Conn]bool),
			broadcast: make(chan message),
		}
		chatRooms[roomID] = chatRoom
		go chatRoom.handleMessages(s.storeMessage)
	}
	roomUnlock(roomID)

	// Cretea upgrade websocket connection
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024 * 1024 * 1024,
		WriteBufferSize: 1024 * 1024 * 1024,
		//Solving cross-domain problems
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Upgrade the HTTP connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go processConn(chatRoom, conn)
}

func processConn(chatRoom *room, conn *websocket.Conn) {
	// Register client connection
	chatRoom.clients[conn] = true
	// Loop to handle incoming messages
	for {
		var msg message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		chatRoom.broadcast <- msg
	}
}

func (cr *room) handleMessages(store func(chatID int64, message message) (int64, error)) {
	for {
		// Get the next message from the chat room's broadcast channel
		msg := <-cr.broadcast
		// Store message in the database
		res, err := store(cr.id, msg)
		if err != nil {
			log.Printf("failed to store message on roomID: %d, err: %v", cr.id, err)
			continue
		}
		// Send the message to all connected clients in the chat room
		for client := range cr.clients {
			err := client.WriteJSON(res)
			if err != nil {
				log.Printf("error sending message to roomID: %d ,error: %v", cr.id, err)
				client.Close()
				delete(cr.clients, client)
			}
		}
	}
}
