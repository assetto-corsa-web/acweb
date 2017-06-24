package user

import (
	"github.com/DeKugelschieber/go-util"
	log "github.com/sirupsen/logrus"
	"model"
)

func RemoveUser(id int64) error {
	user, err := model.GetUserById(id)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading user by ID")
		return util.OpError{1, "Error reading user"}
	}

	if err := user.Remove(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error removing user")
		return util.OpError{2, "Error removing user"}
	}

	return nil
}
