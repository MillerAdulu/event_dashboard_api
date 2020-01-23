package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	userMqtt "github.com/MillerAdulu/dashboard/v1/user/mqtt"
	_userRepository "github.com/MillerAdulu/dashboard/v1/user/repository"
	_userUsecase "github.com/MillerAdulu/dashboard/v1/user/usecase"
	"github.com/centrifugal/gocent"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var (
	centrifuge = gocent.New(gocent.Config{
		Addr: "http://localhost:9099",
		Key:  "6909b7db-9b10-46ae-a826-51618eb61b85",
	})

	tOut = time.Duration(10) * time.Second

	mqttURI, _ = url.Parse("http://134.209.20.138:4117/")
)

var wg sync.WaitGroup

func connect(clientID string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token: %v\n", token)
	return client
}

func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	// opts.SetUsername(uri.User.Username())
	// password, _ := uri.User.Password()
	// opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

func listen(client mqtt.Client, topic string, handler mqtt.MessageHandler) {
	client.Subscribe(topic, 0, handler)
	wg.Done()
}

func rethinkConnect() *r.Session {
	s, err := r.Connect(r.ConnectOpts{
		Address:  "",
		Database: "",
		Username: "",
		Password: "",
		AuthKey:  "",
	})

	if err != nil {
		log.Printf("Query not....: %v", err)
	}
	return s
}

func main() {
	// Clients

	// MQTT Client
	mC := connect("dashboard_golang", mqttURI)

	// Rethink Connection
	rC := rethinkConnect()

	// Repositories
	userRepo := _userRepository.NewUserRepository(rC)

	// Usecases
	allyUsecase := _userUsecase.NewUsecase(centrifuge, userRepo, tOut)

	// Handlers
	_aDel := userMqtt.NewDelivery(&mC, allyUsecase)

	// MQTT Subscriptions
	wg.Add(1)
	go listen(mC, "topics/client/57475/presence", _aDel.Presence)

	wg.Add(1)
	go listen(mC, "topics/client/57475/registration", _aDel.Registration)

	wg.Wait()

	for {
	}
}
