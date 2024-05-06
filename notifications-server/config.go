package main

import (
	"fmt"
	"log"

	"github.com/DavidSie/notification-service/pkg/model"
	"github.com/spf13/viper"
	simplemail "github.com/xhit/go-simple-mail/v2"
)

func populateConfig(app *model.AppConfig) {
	viper.SetConfigName("config")                      // name of config file (without extension)
	viper.SetConfigType("yaml")                        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/notification-service/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.notification-service") // call multiple times to add many search paths
	viper.AddConfigPath(".")                           // optionally look for config in the working directory
	err := viper.ReadInConfig()                        // Find and read the config file
	if err != nil {                                    // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// SMTP config
	app.Stmp.Host = viper.GetString("smtp.host")
	viper.SetDefault("smtp.port", 25)
	app.Stmp.Port = viper.GetInt("smtp.port")
	app.Stmp.KeepAlive = viper.GetBool("smtp.keep_alive")
	app.Stmp.Username = viper.GetString("smtp.user")
	app.Stmp.Password = viper.GetString("smtp.password")
	switch smtpAuth := viper.GetString("smtp.auth_type"); smtpAuth {
	case "Plain":
		app.Stmp.Authentication = simplemail.AuthPlain
	case "Login":
		app.Stmp.Authentication = simplemail.AuthLogin
	case "CRAMMD5":
		app.Stmp.Authentication = simplemail.AuthCRAMMD5
	case "None":
		app.Stmp.Authentication = simplemail.AuthNone
	case "Auto":
		app.Stmp.Authentication = simplemail.AuthAuto
	default:
		log.Printf("No SMTP auth method choosen, valid options are: Plain, Login, CRAMMD5, None, Auto")
	}
	//app.Stmp.Authentication = mail.AuthType(viper.GetInt("smtp.auth_type"))
	switch smtpEncryption := viper.GetString("smtp.encryption"); smtpEncryption {
	case "None":
		app.Stmp.Encryption = simplemail.EncryptionNone
	case "SSL":
		app.Stmp.Encryption = simplemail.EncryptionSSL
	case "TLS":
		app.Stmp.Encryption = simplemail.EncryptionTLS
	case "SSL/TLS":
		app.Stmp.Encryption = simplemail.EncryptionSSLTLS
	case "STARTTLS":
		app.Stmp.Encryption = simplemail.EncryptionSTARTTLS
	default:
		log.Printf("No SMTP encryption method choosen, valid options are: None, SSL, TLS, SSL/TLS, STARTTLS")
	}
	viper.SetDefault("smtp.connect_timeout", "15s")
	app.Stmp.ConnectTimeout = viper.GetDuration("smtp.connect_timeout")
	viper.SetDefault("smtp.send_timeout", "15s")
	app.Stmp.SendTimeout = viper.GetDuration("smtp.send_timeout")

	// Kafka config
	app.Kafka.BootstrapServers = viper.GetString("kafka.bootstrap_servers")
	app.Kafka.SecurityProtocol = viper.GetString("kafka.security_protocol")
	app.Kafka.GroupID = viper.GetString("kafka.group_id")
	app.Kafka.GoApplicationRebalanceEnable = viper.GetBool("kafka.go_application_rebalance_enable")
	log.Printf("\n==============================\nSMTP Configuration:\n SMTP server: %s:%v\n Login: %s\n Auth: %s\n Encryption: %s\n==============================", app.Stmp.Host, app.Stmp.Port, app.Stmp.Username, app.Stmp.Authentication, app.Stmp.Encryption)
	log.Printf("\n==============================\nKafka Configuration:\n Bootstrap Server:%s\n GoApplicationRebalanceEnable: %v\n GroupID: %s\n Security Protocol: %s", app.Kafka.BootstrapServers, app.Kafka.GoApplicationRebalanceEnable, app.Kafka.GroupID, app.Kafka.SecurityProtocol)
}
