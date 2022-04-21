package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	models "github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	sc, _ := stan.Connect(viper.GetString("nats.clusterID"), viper.GetString("nats.clientProducer"), stan.NatsURL(viper.GetString("nats.serverURL")))
	defer sc.Close()

	jsonFile, err := os.Open(viper.GetString("model.jsonFilePath"))
	if err != nil {
		logrus.Fatalf("error open json file: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var model models.Model
	json.Unmarshal(byteValue, &model)
	fmt.Println(model)

	sc.Publish(viper.GetString("nats.subject"), byteValue)
	
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}