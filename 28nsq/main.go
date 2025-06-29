package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/nsqio/go-nsq"
)

type myHandler struct{}

func (*myHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}

func main() {
	var nsqtopic = "learningnsq"

	config := nsq.NewConfig()
	w, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic("Error while connecting producer to nsqd")
	}

	for i := 0; i < 100; i++ {
		err := w.Publish(nsqtopic, []byte("implementing nsq with golang step: "+strconv.Itoa(i)))
		if err != nil {
			log.Panic("Publishing of message on topic : " + nsqtopic + " failed")
		}
	}

	w.Stop()

	var wg sync.WaitGroup
	wg.Add(1)

	config = nsq.NewConfig()
	q, err := nsq.NewConsumer(nsqtopic, "channel", config)
	if err != nil {
		log.Panic("Error while connecting consumer to nsqd")
	}

	//set the handler to the consumer to receive and process message
	q.AddHandler(&myHandler{})

	//connect to my nsqd
	err1 := q.ConnectToNSQD("127.0.0.1:4150")
	if err1 != nil {
		log.Panic("Could not connect to nsqd")
	}

	wg.Wait()
}
