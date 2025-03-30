package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	firstChannel := make(chan Message)
	secondChannel := make(chan Message)
	var msgId int64 = 0
	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&msgId, 1)
			msg := Message{msgId, "RabbitMQ Data"}
			firstChannel <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&msgId, 1)
			msg := Message{msgId, "Apache Kafka Data"}
			secondChannel <- msg
		}
	}()

	//Neste caso apos receber firstChannel e secondChannel apenas vai dar timeout porque os canais não vão encher novamente
	for {
		select {
		case msg := <-firstChannel:
			fmt.Printf("message:id:%v, value:%s\n", msg.id, msg.Msg)

		case msg := <-secondChannel:
			fmt.Printf("message:id:%v, value:%s\n", msg.id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}

}
