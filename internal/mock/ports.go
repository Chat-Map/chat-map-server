// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/application/ports.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	core "github.com/Chat-Map/chat-map-server/internal/core"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetAllUsers mocks base method.
func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]core.UserBySearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]core.UserBySearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserRepositoryMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserRepository)(nil).GetAllUsers), ctx)
}

// GetByEmail mocks base method.
func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserRepositoryMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetByEmail), ctx, email)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(ctx context.Context, userID int64) (core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, userID)
	ret0, _ := ret[0].(core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), ctx, userID)
}

// SearchUserByAll mocks base method.
func (m *MockUserRepository) SearchUserByAll(ctx context.Context, pattern string) ([]core.UserBySearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUserByAll", ctx, pattern)
	ret0, _ := ret[0].([]core.UserBySearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUserByAll indicates an expected call of SearchUserByAll.
func (mr *MockUserRepositoryMockRecorder) SearchUserByAll(ctx, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUserByAll", reflect.TypeOf((*MockUserRepository)(nil).SearchUserByAll), ctx, pattern)
}

// StoreUser mocks base method.
func (m *MockUserRepository) StoreUser(ctx context.Context, user core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUser indicates an expected call of StoreUser.
func (mr *MockUserRepositoryMockRecorder) StoreUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUser", reflect.TypeOf((*MockUserRepository)(nil).StoreUser), ctx, user)
}

// MockChatRepository is a mock of ChatRepository interface.
type MockChatRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatRepositoryMockRecorder
}

// MockChatRepositoryMockRecorder is the mock recorder for MockChatRepository.
type MockChatRepositoryMockRecorder struct {
	mock *MockChatRepository
}

// NewMockChatRepository creates a new mock instance.
func NewMockChatRepository(ctrl *gomock.Controller) *MockChatRepository {
	mock := &MockChatRepository{ctrl: ctrl}
	mock.recorder = &MockChatRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatRepository) EXPECT() *MockChatRepositoryMockRecorder {
	return m.recorder
}

// CreatePrivateChat mocks base method.
func (m *MockChatRepository) CreatePrivateChat(ctx context.Context, userIDs []int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePrivateChat", ctx, userIDs)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePrivateChat indicates an expected call of CreatePrivateChat.
func (mr *MockChatRepositoryMockRecorder) CreatePrivateChat(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrivateChat", reflect.TypeOf((*MockChatRepository)(nil).CreatePrivateChat), ctx, userIDs)
}

// GetChat mocks base method.
func (m *MockChatRepository) GetChat(ctx context.Context, chatID int64) (core.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChat", ctx, chatID)
	ret0, _ := ret[0].(core.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChat indicates an expected call of GetChat.
func (mr *MockChatRepositoryMockRecorder) GetChat(ctx, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChat", reflect.TypeOf((*MockChatRepository)(nil).GetChat), ctx, chatID)
}

// GetChatsMetadata mocks base method.
func (m *MockChatRepository) GetChatsMetadata(ctx context.Context, userID int64) ([]core.ChatMetaData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatsMetadata", ctx, userID)
	ret0, _ := ret[0].([]core.ChatMetaData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatsMetadata indicates an expected call of GetChatsMetadata.
func (mr *MockChatRepositoryMockRecorder) GetChatsMetadata(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatsMetadata", reflect.TypeOf((*MockChatRepository)(nil).GetChatsMetadata), ctx, userID)
}

// IsChatMember mocks base method.
func (m *MockChatRepository) IsChatMember(ctx context.Context, chatID, userID int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsChatMember", ctx, chatID, userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsChatMember indicates an expected call of IsChatMember.
func (mr *MockChatRepositoryMockRecorder) IsChatMember(ctx, chatID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsChatMember", reflect.TypeOf((*MockChatRepository)(nil).IsChatMember), ctx, chatID, userID)
}

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// StoreMessage mocks base method.
func (m *MockMessageRepository) StoreMessage(ctx context.Context, chatID int64, message core.Message) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreMessage", ctx, chatID, message)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreMessage indicates an expected call of StoreMessage.
func (mr *MockMessageRepositoryMockRecorder) StoreMessage(ctx, chatID, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMessage", reflect.TypeOf((*MockMessageRepository)(nil).StoreMessage), ctx, chatID, message)
}

// MockSessionsRepository is a mock of SessionsRepository interface.
type MockSessionsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionsRepositoryMockRecorder
}

// MockSessionsRepositoryMockRecorder is the mock recorder for MockSessionsRepository.
type MockSessionsRepositoryMockRecorder struct {
	mock *MockSessionsRepository
}

// NewMockSessionsRepository creates a new mock instance.
func NewMockSessionsRepository(ctrl *gomock.Controller) *MockSessionsRepository {
	mock := &MockSessionsRepository{ctrl: ctrl}
	mock.recorder = &MockSessionsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionsRepository) EXPECT() *MockSessionsRepositoryMockRecorder {
	return m.recorder
}

// GetSession mocks base method.
func (m *MockSessionsRepository) GetSession(ctx context.Context, sessionID uuid.UUID) (core.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, sessionID)
	ret0, _ := ret[0].(core.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockSessionsRepositoryMockRecorder) GetSession(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockSessionsRepository)(nil).GetSession), ctx, sessionID)
}

// StoreSession mocks base method.
func (m *MockSessionsRepository) StoreSession(ctx context.Context, session core.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSession", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSession indicates an expected call of StoreSession.
func (mr *MockSessionsRepositoryMockRecorder) StoreSession(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSession", reflect.TypeOf((*MockSessionsRepository)(nil).StoreSession), ctx, session)
}

// MockPasswordHasher is a mock of PasswordHasher interface.
type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher.
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance.
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockPasswordHasher) Compare(ctx context.Context, hash, password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", ctx, hash, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Compare indicates an expected call of Compare.
func (mr *MockPasswordHasherMockRecorder) Compare(ctx, hash, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockPasswordHasher)(nil).Compare), ctx, hash, password)
}

// Hash mocks base method.
func (m *MockPasswordHasher) Hash(ctx context.Context, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", ctx, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockPasswordHasherMockRecorder) Hash(ctx, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockPasswordHasher)(nil).Hash), ctx, password)
}

// MockTokenizer is a mock of Tokenizer interface.
type MockTokenizer struct {
	ctrl     *gomock.Controller
	recorder *MockTokenizerMockRecorder
}

// MockTokenizerMockRecorder is the mock recorder for MockTokenizer.
type MockTokenizerMockRecorder struct {
	mock *MockTokenizer
}

// NewMockTokenizer creates a new mock instance.
func NewMockTokenizer(ctrl *gomock.Controller) *MockTokenizer {
	mock := &MockTokenizer{ctrl: ctrl}
	mock.recorder = &MockTokenizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenizer) EXPECT() *MockTokenizerMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenizer) GenerateToken(ctx context.Context, payload core.Payload) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenizerMockRecorder) GenerateToken(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenizer)(nil).GenerateToken), ctx, payload)
}

// ValidateToken mocks base method.
func (m *MockTokenizer) ValidateToken(ctx context.Context, token string) (core.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", ctx, token)
	ret0, _ := ret[0].(core.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockTokenizerMockRecorder) ValidateToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockTokenizer)(nil).ValidateToken), ctx, token)
}

// MockChatNotifier is a mock of ChatNotifier interface.
type MockChatNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockChatNotifierMockRecorder
}

// MockChatNotifierMockRecorder is the mock recorder for MockChatNotifier.
type MockChatNotifierMockRecorder struct {
	mock *MockChatNotifier
}

// NewMockChatNotifier creates a new mock instance.
func NewMockChatNotifier(ctrl *gomock.Controller) *MockChatNotifier {
	mock := &MockChatNotifier{ctrl: ctrl}
	mock.recorder = &MockChatNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatNotifier) EXPECT() *MockChatNotifierMockRecorder {
	return m.recorder
}

// Listen mocks base method.
func (m *MockChatNotifier) Listen(ctx context.Context, address string) chan core.NotifyChat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen", ctx, address)
	ret0, _ := ret[0].(chan core.NotifyChat)
	return ret0
}

// Listen indicates an expected call of Listen.
func (mr *MockChatNotifierMockRecorder) Listen(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockChatNotifier)(nil).Listen), ctx, address)
}

// Notify mocks base method.
func (m *MockChatNotifier) Notify(ctx context.Context, userIDs []int64, chatID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", ctx, userIDs, chatID)
}

// Notify indicates an expected call of Notify.
func (mr *MockChatNotifierMockRecorder) Notify(ctx, userIDs, chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockChatNotifier)(nil).Notify), ctx, userIDs, chatID)
}

// Register mocks base method.
func (m *MockChatNotifier) Register(ctx context.Context, userID int64, address string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", ctx, userID, address)
}

// Register indicates an expected call of Register.
func (mr *MockChatNotifierMockRecorder) Register(ctx, userID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockChatNotifier)(nil).Register), ctx, userID, address)
}

// Unregister mocks base method.
func (m *MockChatNotifier) Unregister(ctx context.Context, userID int64, address string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unregister", ctx, userID, address)
}

// Unregister indicates an expected call of Unregister.
func (mr *MockChatNotifierMockRecorder) Unregister(ctx, userID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockChatNotifier)(nil).Unregister), ctx, userID, address)
}

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidator) Validate(ctx context.Context, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorMockRecorder) Validate(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidator)(nil).Validate), ctx, data)
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockServer) Run(port string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", port)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockServerMockRecorder) Run(port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockServer)(nil).Run), port)
}
