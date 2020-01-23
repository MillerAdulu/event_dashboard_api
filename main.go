package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	allyMqtt "github.com/MillerAdulu/dashboard/v1/ally/mqtt"
	_allyUsecase "github.com/MillerAdulu/dashboard/v1/ally/usecase"
	"github.com/centrifugal/gocent"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

func main() {
	// Clients

	// MQTT Client
	mC := connect("dashboard", mqttURI)

	// Repositories

	// Usecases
	allyUsecase := _allyUsecase.NewUsecase(centrifuge, tOut)

	// Deliveries
	_aDel := allyMqtt.NewDelivery(&mC, allyUsecase)

	// MQTT Subscriptions
	wg.Add(1)
	go listen(mC, "topics/client/57475/presence", _aDel.Presence)

	wg.Add(1)
	go listen(mC, "topics/client/57475", _aDel.Global)

	wg.Wait()

	for {
	}
}
