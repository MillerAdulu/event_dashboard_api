package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/MillerAdulu/dashboard/entities"

	"github.com/MillerAdulu/dashboard/v1/ally"
	"github.com/centrifugal/gocent"
	ms "github.com/mitchellh/mapstructure"
)

type allyUcase struct {
	Cent *gocent.Client
	tOut time.Duration
}

// NewUsecase -
func NewUsecase(cent *gocent.Client, t time.Duration) ally.Usecase {
	return &allyUcase{
		Cent: cent,
		tOut: t,
	}
}

// RegisterAlly - Update statistics for user registrations
func (aU *allyUcase) RegisterAlly(ctx context.Context, ally map[string]interface{}) {
	var err error
	var a entities.UserRegistration
	ch := "registration"

	ctx, cancel := context.WithTimeout(ctx, aU.tOut)
	defer cancel()

	ms.Decode(ally, &a)

	// TODO: Insert data in `a` to RethinkDB

	// Publish data to centrifuge
	data, _ := json.Marshal(a)

	if err != nil {
		log.Fatalf("Error marshaling:  %v", a)
	}

	err = aU.Cent.Publish(ctx, ch, data)

	if err != nil {
		log.Fatalf("Error calling publish: %v", err)
	}

	log.Printf("Successfully published to: %v", ch)

}
