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
	st := &Stream{
		sc: sc,
		services: services,
	}
	st.Streaming()
	return st
}

func (s *Stream) Streaming() {
	s.getModel()
}