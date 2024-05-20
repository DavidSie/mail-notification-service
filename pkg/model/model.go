package model

import (
	"time"

	"crypto/tls"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	mail "github.com/xhit/go-simple-mail/v2"
	models "github.com/DavidSie/go-models/pkg/model"
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
	MailChannel chan models.EmailRequest `yaml:"mail_channel"`
}


type NotificationServerConfig struct {
	KafkaConfigMap kafka.ConfigMap
}
