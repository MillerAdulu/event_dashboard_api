package repository

import (
	"log"

	"github.com/MillerAdulu/dashboard/entities"
	"github.com/MillerAdulu/dashboard/v1/user"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type userRepository struct {
	Session *r.Session
}

// NewUserRepository -
func NewUserRepository(Session *r.Session) user.Repository {
	return &userRepository{Session}
}

func (uR *userRepository) RegisterUser(reg entities.UserRegistrationData) {
	_, err := r.Table("user").Insert(&reg).Run(uR.Session)
	if err != nil {
		log.Printf("Query not....: %v", err)
	}

	log.Println("Success")

}
