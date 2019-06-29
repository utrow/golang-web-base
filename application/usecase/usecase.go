package usecase

type Interacter struct {
	Ping Ping
}

func NewInteractor() Interacter {
	return Interacter{
		Ping: NewPingUsecase(),
	}
}
