package user

import (
	"github.com/MillerAdulu/dashboard/entities"
	"context"
)

// Usecase -
type Usecase interface {
	RegisterUser(ctx context.Context, ally entities.UserRegistrationData)
	DeleteUser(ctx context.Context, userID int)
}
