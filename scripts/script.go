package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

const (
	modelsPath = "./data"
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

	const channelName = "foo"

	ticker := time.NewTicker(5 * time.Second)

	fileNumber := 1
	for range ticker.C {
		if fileNumber == 6 {
			break
		}
		fullPath := fmt.Sprintf("%s/test%d.json", modelsPath, fileNumber)
		bytes, err := parseFile(fullPath)
		if err != nil {
			return
		}

		err = sc.Publish(channelName, bytes)
		if err != nil {
			log.Printf("Ошибка публикации сообщения: %v\n", err.Error())
		} else {
			log.Printf("Отправлено сообщение: %s\n", bytes)
		}
		fileNumber++
	}
	defer ticker.Stop()
}

func parseFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read file: %s: %s\n", path, err.Error())
		return nil, err
	}

	bytes, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("Failed to marshal: %s\n", err.Error())
		return nil, err
	}

	return bytes, err
}
