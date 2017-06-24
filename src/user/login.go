package user

import (
	"github.com/DeKugelschieber/go-util"
	log "github.com/sirupsen/logrus"
	"model"
)

func Login(loginEmail, pwd string) (*model.User, error) {
	loginEmail = util.Trim(loginEmail)
	pwd = util.Trim(pwd)

	if loginEmail == "" || pwd == "" {
		return nil, util.OpError{1, "Login and password must be set"}
	}

	// read user
	pwd = util.Sha256base64(pwd)
	user, err := model.GetUserByLoginOrEmailAndPassword(loginEmail, loginEmail, pwd)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("User could not be found on login")
		return nil, util.OpError{2, "User not found"}
	}

	// session is started within the HTTP handler
	return user, nil
}
