package stream

import (
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/service"
	"github.com/nats-io/stan.go"
)

type Stream struct {
	sc stan.Conn
	services *service.Service
}

func NewStream(sc stan.Conn, services *service.Service) *Stream {
	return &Stream{
		sc: sc,
		services: services,
	}
}

func Streaming(s *Stream) {
	s.getModel()
}