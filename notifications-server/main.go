package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/DavidSie/notification-service/pkg/mail"
	"github.com/DavidSie/notification-service/pkg/model"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var wg sync.WaitGroup
var app model.AppConfig

func main() {

	populateConfig(&app)
	nsc := model.NotificationServerConfig{KafkaConfigMap: kafka.ConfigMap{
		"bootstrap.servers":               app.Kafka.BootstrapServers,
		"security.protocol":               app.Kafka.SecurityProtocol,
		"group.id":                        app.Kafka.GroupID,
		"go.application.rebalance.enable": app.Kafka.GoApplicationRebalanceEnable},
	}
	app.MailChannel = make(chan model.EmailRequest)
	defer close(app.MailChannel)
	mailSrv := mail.Mailer{AppConfig: app}
	mailSrv.ListenForMail()

	// for each notification channel create on go routine
	wg.Add(1)
	go HandleEmailRequests(&nsc.KafkaConfigMap, app)
	wg.Wait()
}

func HandleEmailRequests(kcm *kafka.ConfigMap, app model.AppConfig) {
	defer wg.Done()
	consumer, err := kafka.NewConsumer(kcm)
	defer func() {
		err = consumer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	topics := []string{model.RequestEmailTopic}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatal(err)
	}
	run := true

	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			emailRequest := model.EmailRequest{}
			err := json.Unmarshal(e.Value, &emailRequest)
			if err != nil {
				fmt.Printf("Error while unmarshaling kafka message to emailRequest %v", err)
			}
			app.MailChannel <- emailRequest
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
			// default:
			// 	fmt.Printf("Ignored %v\n", e)
		}
	}

}
