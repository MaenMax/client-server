package main

import (
	queue2 "client-server/src/queue"
	"client-server/src/queue/subscriber"
	l4g "github.com/alecthomas/log4go"
)

var (
	LogFile string = "conf/server.xml"
)

func main() {
	l4g.LoadConfiguration(LogFile)
	l4g.Info("Server is starting..")
	go subscriber.Run() //Starting subscriber go routine

	var msg_queue = queue2.New()

	//server's main for loop

	for {
		select {
		case msg := <-subscriber.RecvChan:
			l4g.Info("Received a message: %s\n", string(msg.Data))
			err := subscriber.ProcessMessage(msg.Data, msg_queue)
			if err != nil {
				continue
			}
			msg_queue.Print()
		}
	}

	l4g.Info("Server stopped!")
}
