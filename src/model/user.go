package model

import (
	"errors"
)

type User struct {
	Id        int64  `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Pwd       string `db:"password"` // do not send this to client
	Admin     bool   `json:"admin"`
	Moderator bool   `json:"moderator"`
}

func (m *User) Save() error {
	if m.Id == 0 {
		res, err := session.NamedExec("INSERT INTO user (login, email, password, admin, moderator) VALUES (:login, :email, :password, :admin, :moderator)", m)

		if err != nil {
			return err
		}

		id, err := res.LastInsertId()

		if err != nil {
			return err
		}

		m.Id = id
		return nil
	}

	_, err := session.NamedExec("UPDATE user SET login = :login, email = :email, password = :password, admin = :admin, moderator = :moderator WHERE id = :id", m)
	return err
}

func (m *User) Remove() error {
	if m.Id == 0 {
		return errors.New("ID must be set")
	}

	_, err := session.NamedExec("DELETE FROM user WHERE id = :id", m)
	return err
}

func GetUserByLoginOrEmailAndPassword(login, email, pwd string) (*User, error) {
	user := new(User)

	if err := session.Get(user, "SELECT * FROM user WHERE (login LIKE ? OR email LIKE ?) AND password = ?", login, email, pwd); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(id int64) (*User, error) {
	user := new(User)

	if err := session.Get(user, "SELECT * FROM user WHERE id = ?", id); err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByLoginOrEmail(login, email string) (*User, error) {
	user := new(User)

	if err := session.Get(user, "SELECT * FROM user WHERE login LIKE ? OR email LIKE ?", login, email); err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUser() ([]User, error) {
	user := make([]User, 0)

	if err := session.Select(&user, "SELECT * FROM user ORDER BY login, email ASC"); err != nil {
		return nil, err
	}

	return user, nil
}
