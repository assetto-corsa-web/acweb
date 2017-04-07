package user

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
)

func RemoveUser(id int64) error {
	user, err := model.GetUserById(id)

	if err != nil {
		log.Printf("Error reading user by ID: %v", err)
		return util.OpError{1, "Error reading user"}
	}

	if err := user.Remove(); err != nil {
		log.Printf("Error removing user: %v", err)
		return util.OpError{2, "Error removing user"}
	}

	return nil
}
