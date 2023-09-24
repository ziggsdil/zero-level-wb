package main

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func main() {
	natsURL := "nats://localhost:4222"
	clusterID := "zero-level-server"
	clientID := "zero-level-client-publish"

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к серверу NATS: %v\n", err.Error())
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Ошибка к подключению к серверу NATS Streaming: %v\n", err.Error())
	}
	defer sc.Close()

	channelName := "foo"

	for {
		message := "Hello world!"
		err := sc.Publish(channelName, []byte(message))
		if err != nil {
			log.Printf("Ошибка публикации сообщения: %v\n", err.Error())
		} else {
			log.Printf("Отправлено сообщение: %s\n", message)
		}

		time.Sleep(time.Second * 5)
	}
}
