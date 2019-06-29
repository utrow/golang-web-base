package usecase

import (
	"github.com/utrow/golang-web-base/application/usecase/in"
	"github.com/utrow/golang-web-base/application/usecase/out"
)

type Ping interface {
	Get(ipt *in.GetPing) (*out.GetPingResponse, error)
}

type ping struct {
	ping Ping
}

func NewPingUsecase() Ping {
	return &ping{}
}

func (*ping) Get(ipt *in.GetPing) (*out.GetPingResponse, error) {
	return &out.GetPingResponse{
		Message: ipt.Text,
	}, nil
}
