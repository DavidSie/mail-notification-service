
## Development

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

### Start applications

```
go run ./notifications-request-producer/
go run ./notifications-server/

```