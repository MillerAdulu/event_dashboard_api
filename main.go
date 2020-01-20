package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/MillerAdulu/dashboard/utils/websocket"
	_socketHandler "github.com/MillerAdulu/dashboard/v1/websocket"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := websocket.New("ws://0.0.0.0:0/socket")

	socket.OnConnected = func(socket websocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket websocket.Socket) {
		log.Println("Received connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket websocket.Socket) {
		log.Println("Received message " + message)
		_socketHandler.SocketDataHandler(message)
	}

	socket.OnBinaryMessage = func(data []byte, socket websocket.Socket) {
		log.Println("Received binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket websocket.Socket) {
		log.Println("Received ping " + data)
	}

	socket.OnPongReceived = func(data string, socket websocket.Socket) {
		log.Println("Received pong " + data)
	}

	socket.OnDisconnected = func(err error, socket websocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}

}
