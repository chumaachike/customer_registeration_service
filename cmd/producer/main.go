package main

import (
	"log"
	"time"

	"github.com/chumaachike/eventdirvenrabbitrabbit/internal"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("admin", "admin123", "localhost:5672", "customers")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	if err := client.CreateQueue("customers_created", true, false); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	log.Println(client)
}
