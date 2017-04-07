package user

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
)

func GetAllUser() ([]model.User, error) {
	user, err := model.GetAllUser()

	if err != nil {
		log.Printf("Error reading all user: %v", err)
		return nil, util.OpError{1, "Error reading all user"}
	}

	return user, nil
}

func GetUser(id int64) (*model.User, error) {
	user, err := model.GetUserById(id)

	if err != nil {
		log.Printf("Error reading user by ID: %v", err)
		return nil, util.OpError{1, "Error reading user"}
	}

	return user, nil
}
