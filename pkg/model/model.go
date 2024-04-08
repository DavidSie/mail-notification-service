package model

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

const (
	RequestEmailTopic string = "RequestEmailTopic"
)

type Email string

type EmailRequest struct {
	Recipients    []Email `json:"recipients"`
	CcRecipients  []Email `json:"cc_recipients"`
	BccRecipients []Email `json:"bcc_recipients"`
	Sender        Email   `json:"sender"`
	Title         string  `json:"title"`
	Message       string  `json:"message"`
}

type MailingService interface {
	Send(emailRequest EmailRequest) error
}

type NotificationServerConfig struct {
	KafkaConfigMap kafka.ConfigMap
}
