package user

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
)

func Login(loginEmail, pwd string) (int64, error) {
	loginEmail = util.Trim(loginEmail)
	pwd = util.Trim(pwd)

	if loginEmail == "" || pwd == "" {
		return 0, util.OpError{1, "Login and password must be set"}
	}

	// read user
	pwd = util.Md5base64(pwd)
	user, err := model.GetUserByLoginOrEmailAndPassword(loginEmail, loginEmail, pwd)

	if err != nil {
		log.Printf("User could not be found on login: %v", err)
		return 0, util.OpError{2, "User not found"}
	}

	// session is started within the HTTP handler
	return user.Id, nil
}
