package main

import (
	"flag"
	"fmt"
	nats "github.com/nats-io/nats.go"
	"strings"
)

var (
	message      = flag.String("msg", "", "Data to be published to the queue service")
	NatsServers  = []string{"http://127.0.0.1:4222", "http://127.0.0.1:5222", "http://127.0.0.1:6222"}
	QueueSubject = "bloxroute"
	err          error
)

func main() {
	flag.Parse()
	// Connect to a server
	nc, err := nats.Connect(strings.Join(NatsServers, ","))
	if err != nil {
		fmt.Println("Could not connect to NATS message queue servers. ", err)
	}

	fmt.Println("Publishing ...")

	// Simple Publisher
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		nc.Publish(QueueSubject, []byte(*message))
	//	}()
	//}
	//fmt.Scanln()
	nc.Publish(QueueSubject, []byte(*message))
	fmt.Println("Published message: ", *message)
	fmt.Println("Published!")
	nc.Close()

}
