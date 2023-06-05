package application

type UseCase struct {
	v  Validator
	ph PasswordHasher
	tk Tokenizer

	ur UserRepository
	sr SessionsRepository
	cr ChatRepository
	mr MessageRepository

	Command
}

type Command struct {
	Signin        SigninCommand
	Signup        SignupCommand
	ChatGet       GetChatCommand
	ChatCreate    CreateChatCommand
	ChatMeta      GetChatMetaCommand
	MessageStore  StoreMessageCommand
	ValidateToken TokenValidateCommand
}

type UseCaseOption func(*UseCase)

func NewUseCase(opts ...UseCaseOption) *UseCase {
	uc := &UseCase{}
	for _, opt := range opts {
		opt(uc)
	}
	uc.Command = Command{
		Signin: NewSigninCommandImplV1(uc.v, uc.ur, uc.sr, uc.ph, uc.tk),
		Signup: NewSignupCommandImplV1(uc.v, uc.ur, uc.ph),

		ChatGet:    NewGetChatCommandImplV1(uc.v, uc.cr),
		ChatCreate: NewCreateChatCommandImplV1(uc.v, uc.cr),
		ChatMeta:   NewGetChatMetaCommandImplV1(uc.cr),

		MessageStore: NewStoreMessageCommandImplV1(uc.v, uc.ur, uc.cr, uc.mr),

		ValidateToken: NewTokenValidateCommandImplV1(uc.v, uc.tk),
	}
	return uc
}

func WithUserRepository(ur UserRepository) UseCaseOption {
	return func(uc *UseCase) {
		uc.ur = ur
	}
}

func WithSessionsRepository(sr SessionsRepository) UseCaseOption {
	return func(uc *UseCase) {
		uc.sr = sr
	}
}

func WithChatRepository(cr ChatRepository) UseCaseOption {
	return func(uc *UseCase) {
		uc.cr = cr
	}
}

func WithMessageRepository(mr MessageRepository) UseCaseOption {
	return func(uc *UseCase) {
		uc.mr = mr
	}
}

func WithPasswordHasher(p PasswordHasher) UseCaseOption {
	return func(uc *UseCase) {
		uc.ph = p
	}
}

func WithValidator(v Validator) UseCaseOption {
	return func(uc *UseCase) {
		uc.v = v
	}
}

func WithTokenizer(t Tokenizer) UseCaseOption {
	return func(uc *UseCase) {
		uc.tk = t
	}
}
