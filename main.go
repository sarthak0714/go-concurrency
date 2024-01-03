package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	messagechan chan Message
	quitchan    chan struct{}
}

func (s *Server) StartAndListen() {
free:
	for {
		select {
		case msg := <-s.messagechan:
			fmt.Printf("%s SAYS %s\n", msg.From, msg.Payload)
		case <-s.quitchan:
			fmt.Println("Shutting Down...")
			break free
		default:
		}
	}
	fmt.Println("Server Has Shut Down.")
}

func SendMessageToServer(messagechan chan Message, payload string) {
	msg := Message{
		From:    "Shrihari",
		Payload: payload,
	}
	messagechan <- msg
}

func QuitFullServer(quitchan chan struct{}) {
	close(quitchan)
}

func main() {
	s := &Server{
		messagechan: make(chan Message),
		quitchan:    make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		SendMessageToServer(s.messagechan, "Shreyas says hi!")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		QuitFullServer(s.quitchan)
	}()

	select {}
}
