package application

type UseCase struct {
	ur UserRepository
	sr SessionsRepository
	cr ChatRepository
	mr MessageRepository

	Command
}

type Command struct{}

type UseCaseOption func(*UseCase)

func NewUseCase(opts ...UseCaseOption) *UseCase {
	uc := &UseCase{}
	for _, opt := range opts {
		opt(uc)
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
