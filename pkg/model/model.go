package model

import (
	"time"

	"crypto/tls"

	mail "github.com/xhit/go-simple-mail/v2"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	RequestEmailTopic string = "RequestEmailTopic"
)

type StmpConfig struct {
	Host           string          `yaml:"host"`
	Port           int             `yaml:"port"`
	Authentication mail.AuthType   `yaml:"auth_type"`
	Username       string          `yaml:"username"`
	Password       string          `yaml:"password"`
	TLSConfig      *tls.Config     `yaml:"tls_config"`
	Encryption     mail.Encryption `yaml:"auth_encryption"`
	KeepAlive      bool            `yaml:"keep_alive"`
	ConnectTimeout time.Duration   `yaml:"connect_timeout"`
	SendTimeout    time.Duration   `yaml:"send_timeout"`
}

type KafkaConfig struct {
	BootstrapServers             string
	SecurityProtocol             string
	GroupID                      string
	GoApplicationRebalanceEnable bool
}
type AppConfig struct {
	Stmp        StmpConfig        `yaml:"smtp"`
	Kafka       KafkaConfig       `yaml:"kafka"`
	MailChannel chan EmailRequest `yaml:"mail_channel"`
}

// type EmailRequeststore information regarding requested email
type EmailRequest struct {
	Recipients    []string `yaml:"recipients"`
	CcRecipients  []string `yaml:"cc_recipients"`
	BccRecipients []string `yaml:"bcc_recipients"`
	// If Sender cannot be set iw will be ignored
	Sender             string           `yaml:"sender"`
	Title              string           `yaml:"title"`
	Message            string           `yaml:"message"`
	MessageContentType mail.ContentType `yaml:"message_content_type"`
}

type NotificationServerConfig struct {
	KafkaConfigMap kafka.ConfigMap
}
