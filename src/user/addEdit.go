package user

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
)

func AddEditUser(id int64, login, email, pwd1, pwd2 string) error {
	// check fields
	login = util.Trim(login)
	email = util.Trim(email)

	if login == "" || email == "" {
		return util.OpError{1, "Login and email must be set"}
	}

	if (id == 0 && (pwd1 == "" || pwd2 == "")) || (pwd1 != "" && pwd2 != "" && len(pwd1) < 8) {
		return util.OpError{2, "Password invalid"}
	}

	if pwd1 != pwd2 {
		return util.OpError{3, "Passwords must be equal"}
	}

	// look for existing user
	if user, _ := model.FindUserByLoginOrEmail(login, email); user != nil && user.Id != id {
		return util.OpError{4, "Login and/or email exists already"}
	}

	// create/update new user
	var user *model.User

	if id == 0 {
		user = &model.User{0, login, email, util.Md5base64(pwd1)}
	} else {
		existingUser, err := model.GetUserById(id)

		if err != nil {
			log.Printf("Error reading user by ID: %v", err)
			return util.OpError{5, "Error reading user"}
		}

		existingUser.Login = login
		existingUser.Email = email

		if pwd1 != "" {
			existingUser.Pwd = util.Md5base64(pwd1)
		}

		user = existingUser
	}

	if err := user.Save(); err != nil {
		log.Printf("Error saving new user: %v", err)
		return util.OpError{6, "Error saving new user"}
	}

	return nil
}
