package usecase

import (
	"github.com/utrow/golang-web-base/application/usecase/in"
	"github.com/utrow/golang-web-base/application/usecase/out"
	"github.com/utrow/golang-web-base/domain/repository"
)

type Ping interface {
	Get(ipt *in.GetPing) (*out.GetPingResponse, error)
}

type ping struct {
	ping       Ping
	repository repository.Repository
}

func NewPingUsecase(r repository.Repository) Ping {
	return &ping{
		repository: r,
	}
}

func (p *ping) Get(ipt *in.GetPing) (*out.GetPingResponse, error) {
	return &out.GetPingResponse{
		Message: ipt.Text,
	}, nil
}
