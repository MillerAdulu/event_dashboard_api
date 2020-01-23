package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/MillerAdulu/dashboard/entities"

	"github.com/MillerAdulu/dashboard/v1/user"
	"github.com/centrifugal/gocent"
)

type allyUcase struct {
	Cent *gocent.Client
	tOut time.Duration
}

// NewUsecase -
func NewUsecase(cent *gocent.Client, t time.Duration) user.Usecase {
	return &allyUcase{
		Cent: cent,
		tOut: t,
	}
}

// RegisterAlly - Update statistics for user registrations
func (aU *allyUcase) RegisterUser(ctx context.Context, user entities.UserRegistrationData) {
	var err error
	ch := "registration"

	ctx, cancel := context.WithTimeout(ctx, aU.tOut)
	defer cancel()

	// TODO: Insert data in `a` to RethinkDB

	// Publish data to centrifuge
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshaling:  %v", user)
	}

	err = aU.Cent.Publish(ctx, ch, data)

	if err != nil {
		log.Printf("Error calling publish: %v", err)
	}

	log.Printf("Successfully published to: %v", ch)

}

func (aU *allyUcase) DeleteUser(ctx context.Context, userID int) {}
