package user

import (
	"context"
	"github.com/MillerAdulu/dashboard/entities"
)

// Usecase -
type Usecase interface {
	RegisterUser(ctx context.Context, ally entities.UserRegistrationData)
	DeleteUser(ctx context.Context, userID int)
}
