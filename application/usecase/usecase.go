package usecase

import "github.com/utrow/golang-web-base/domain/repository"

type Interacter struct {
	Ping Ping
}

func NewInteractor(repo repository.Repository) Interacter {
	return Interacter{
		Ping: NewPingUsecase(repo),
	}
}
