package stream

import (
	"encoding/json"

	models "github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (s *Stream) getModel() {
	s.sc.Subscribe(viper.GetString("nats.subject"), func(msg *stan.Msg) {
		var model models.Model
		err := json.Unmarshal(msg.Data, &model)
		if err != nil {
			logrus.Errorf("error parse model from stream: %s", err.Error())
			return
		}
		if model.Order_uid == "" {
			logrus.Error("error order_uid cannot be empty")
			return
		}
		_, err = s.services.Create(model)
		if err != nil {
			logrus.Errorf("cannot be create model: %s", err.Error())
			return
		}
	}, stan.DurableName(viper.GetString("nats.clientSubscriber")))
}