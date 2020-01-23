package mqtt

import (
	"fmt"

	"github.com/MillerAdulu/dashboard/v1/ally"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// AllyHandler -
type AllyHandler struct {
	Client *mqtt.Client
	AUcase ally.Usecase
}

// NewDelivery -
func NewDelivery(c *mqtt.Client, a ally.Usecase) AllyHandler {
	return AllyHandler{
		Client: c,
		AUcase: a,
	}
}

// Presence -
func (aM *AllyHandler) Presence(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
}

// Global -
func (aM *AllyHandler) Global(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
}
