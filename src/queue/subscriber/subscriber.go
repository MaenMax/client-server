package subscriber

import (
	queue2 "client-server/src/queue"
	"encoding/json"
	"errors"
	l4g "github.com/alecthomas/log4go"
	nats "github.com/nats-io/nats.go"
	"strings"
)

var (
	NatsServers  = []string{"http://127.0.0.1:4222", "http://127.0.0.1:5222", "http://127.0.0.1:6222"}
	QueueSubject = "bloxroute"
	err          error
	RecvChan     = make(chan *nats.Msg, 64)
)

func Run() {

	// Connect to a server
	nc, err := nats.Connect(strings.Join(NatsServers, ","))
	if err != nil {
		l4g.Error("Could not connect to NATS message queue server. %s", err)
	}

	l4g.Info("Connected to NATS servers.")

	l4g.Info("Listening on '%s' subject..", QueueSubject)

	// Simple Async Subscriber

	//nc.Subscribe(QueueSubject, func(m *nats.Msg) {
	//	l4g.Info("sending to channel")
	//	RecvChan <- m
	//})

	_, err = nc.ChanSubscribe(QueueSubject, RecvChan)
	if err != nil {
		l4g.Error("Unable to subscribe: %s", err)
	}
}

func ProcessMessage(message []byte, queue *queue2.T_Queue) (err error) {
	l4g.Info("Processing Message ...")
	request := &queue2.T_Request{}
	err = json.Unmarshal(message, request)
	if err != nil {
		l4g.Error("Issue with message format!. %s", err)
		return
	}
	switch request.Type {
	case 1: //Add item
		l4g.Info("Received add item request")
		queue.AddItem(request.Item, request.ClientId)
	case 2: //Remove item
		l4g.Info("Received remove item request")
		removed := queue.RemoveItem(request.ClientId, request.Item)
		if !removed {
			err = errors.New("Item not found!")
			return
		}
	case 3: //Get item
		l4g.Info("Received get item request")
		item_data := queue.GetItem(request.Item, request.ClientId)
		if item_data == "" {
			err = errors.New("Item not found!")
			return
		}
		l4g.Info(item_data)
	case 4: //Get all items
		l4g.Info("Received get all items request")
		all_items := queue.GetAllItems(request.Item, request.ClientId)
		if len(all_items) == 0 {
			err = errors.New("No items found for this client.")
			return
		}
		l4g.Info(all_items)
	default:
		err = errors.New("Undefined request type")
		return
	}
	return
}
