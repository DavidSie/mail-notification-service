# syntax=docker/dockerfile:1
FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /bin/notifications-request-producer ./notifications-request-producer/
RUN go build -o /bin/notifications-server ./notifications-server/


FROM ubuntu:22.04
COPY --from=build /bin/notifications-request-producer /bin/notifications-request-producer
COPY --from=build /bin/notifications-server /bin/notifications-server
ENTRYPOINT ["/bin/notifications-server"]
