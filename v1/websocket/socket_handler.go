package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_usersCase "github.com/MillerAdulu/dashboard/v1/ally/usecase"
	"github.com/centrifugal/gocent"
)

var (
	centrifuge = gocent.New(gocent.Config{
		Addr: "http://localhost:9099",
		Key:  "6909b7db-9b10-46ae-a826-51618eb61b85",
	})

	tOut = time.Duration(10) * time.Second

	users = _usersCase.NewUsecase(centrifuge, tOut)
)

// SocketDataHandler - Decode message string and call relevant usecase function
func SocketDataHandler(message string) {
	var msg map[string]interface{}
	var err error
	ctx := context.Background()

	err = json.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
	}

	// Conditional check for the message events and call to relevant methods
	switch msg["communication_event"] {
	case "ON_USER_REGISTRATION":
		fmt.Printf("Communication event 1: %v", msg)
		users.RegisterAlly(ctx, msg["data"].(map[string]interface{}))
	case "ON_USER_1":
		fmt.Printf("Communication event 2: %v", msg)
	case "ON_USER_2":
		fmt.Printf("Communication event 3: %v", msg)
	default:
		fmt.Printf("Undefined communication event: %v", msg)
	}

}
