package stream

import (
	"fmt"

	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
)

func (s *Stream) getModel() {
	s.sc.Subscribe(viper.GetString("nats.subject"), func(msg *stan.Msg) {
		fmt.Println(msg.Data)
	}, stan.DurableName(viper.GetString("nats.clientSubscriber")))
	//fmt.Println()
}