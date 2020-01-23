package user

import (
	"github.com/MillerAdulu/dashboard/entities"
)

// Repository -
type Repository interface {
	RegisterUser(reg entities.UserRegistrationData)
}
