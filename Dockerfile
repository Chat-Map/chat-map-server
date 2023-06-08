FROM golang:1.19 AS development
WORKDIR /go/src/github.com/chatmap/server
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cespare/reflex@latest
CMD reflex -sr '\.go$' go run ./cmd/main.go

FROM golang:alpine AS builder
WORKDIR /go/src/github.com/chatmap/server
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY go.mod go.sum ./
RUN go build -o /go/bin/server ./cmd/main.go

FROM alpine:latest AS production
COPY --from=builder /go/bin/server /go/bin/server
COPY ./docs /docs
ENTRYPOINT ["/go/bin/server"]
