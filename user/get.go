package user

import (
	log "github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/acweb/model"
	"github.com/assetto-corsa-web/acweb/util"
)

func GetAllUser() ([]model.User, error) {
	user, err := model.GetAllUser()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading all user")
		return nil, util.OpError{1, "Error reading all user"}
	}

	return user, nil
}

func GetUser(id int64) (*model.User, error) {
	user, err := model.GetUserById(id)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading user by ID")
		return nil, util.OpError{1, "Error reading user"}
	}

	return user, nil
}
