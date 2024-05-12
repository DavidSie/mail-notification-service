package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/DavidSie/notification-service/pkg/model"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	mail "github.com/xhit/go-simple-mail/v2"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"security.protocol": "plaintext",
		"client.id":         "Test client",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	defer p.Close()

	wb := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wb.Add(1)
		go func(iterator int) {
			defer wb.Done()

			emailRequest := model.EmailRequest{
				Recipients:         []string{"someone@mail.com"},
				Sender:             "iam.sender@email.com",
				Title:              fmt.Sprintf("Message number: %d", iterator),
				Message:            fmt.Sprintf("<strong>Message number: %d</strong><br>Dear recipient, I'm glad you got your mail number %d", iterator, iterator),
				MessageContentType: mail.TextHTML,
			}

			r := rand.Intn(60)
			time.Sleep(time.Duration(r) * time.Second)
			err = RequestMailSending(emailRequest, p)
			if err != nil {
				fmt.Printf("Error while requesting email sending: %v\n", err)
			}
		}(i)
	}
	wb.Wait()

}

func RequestMailSending(emailRequest model.EmailRequest, p *kafka.Producer) error {

	delivery_chan := make(chan kafka.Event, 10000)

	topic := model.RequestEmailTopic

	fmt.Printf("Requesting email send to: %s \n", emailRequest.Recipients)
	value, err := json.Marshal(emailRequest)
	if err != nil {
		fmt.Printf("Failed to parse Email Request to Byte Slice: %s\n", err)
		return err
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		delivery_chan,
	)
	if err != nil {
		fmt.Printf("Failed to send message: %s\n", err)
		return err
	}
	return nil
}
