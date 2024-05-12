# Mail notification service

This is a microservice aiming at sending mails upon receiving Mail reques via Kafka event.

# Development

In order to run and test localy the service, one needs smtp mail server and kafka, and producer of events.
For SMTP server one can use `mailhog`, to produce events one can use notifications-request-producer.



### Start Kafka container 

```
docker run -p 9092:9092 -d --rm apache/kafka:3.7.0
```


### Test mail server


install 
```
brew update && brew install mailhog
```

start 
```
brew services start mailhog
```

more info: https://github.com/mailhog/MailHog


### Events producer

```bash
go run ./notifications-request-producer/ 
```

### Start notififaction server

```bash

go run ./notifications-server/
```
