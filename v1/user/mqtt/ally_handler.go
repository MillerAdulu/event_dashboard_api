package mqtt

import (
	"context"
	"fmt"

	"github.com/MillerAdulu/dashboard/entities"

	"github.com/MillerAdulu/dashboard/v1/user"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	ms "github.com/mitchellh/mapstructure"
)

// AllyHandler -
type AllyHandler struct {
	Client *mqtt.Client
	AUcase user.Usecase
}

// NewDelivery -
func NewDelivery(c *mqtt.Client, a user.Usecase) AllyHandler {
	return AllyHandler{
		Client: c,
		AUcase: a,
	}
}

// Presence -
func (aM *AllyHandler) Presence(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
}

// Registration -
func (aM *AllyHandler) Registration(c mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	var a entities.UserRegistrationData
	ctx := context.Background()

	ms.Decode(msg.Payload(), &a)

	aM.AUcase.RegisterUser(ctx, a)

}
