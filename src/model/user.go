package model

import (
	"database/sql"
	"db"
	"errors"
	"github.com/DeKugelschieber/go-util"
)

type User struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pwd   string // do not send this to client
}

func (m *User) Save() error {
	if m.Id == 0 {
		res, err := db.Get().Exec("INSERT INTO user (login, email, password) VALUES (?, ?, ?)", m.Login,
			m.Email,
			m.Pwd)

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

	_, err := db.Get().Exec("UPDATE user SET login = ?, email = ?, password = ? WHERE id = ?", m.Login,
		m.Email,
		m.Pwd,
		m.Id)
	return err
}

func (m *User) Remove() error {
	if m.Id == 0 {
		return errors.New("ID must be set")
	}

	_, err := db.Get().Exec("DELETE FROM user WHERE id = ?", m.Id)
	return err
}

func GetUserByLoginOrEmailAndPassword(login, email, pwd string) (*User, error) {
	row := db.Get().QueryRow("SELECT * FROM user WHERE (login LIKE ? OR email LIKE ?) AND password = ?", login, email, pwd)
	return scanOneUser(row)
}

func GetUserById(id int64) (*User, error) {
	row := db.Get().QueryRow("SELECT * FROM user WHERE id = ?", id)
	return scanOneUser(row)
}

func FindUserByLoginOrEmail(login, email string) (*User, error) {
	row := db.Get().QueryRow("SELECT * FROM user WHERE login LIKE ? OR email LIKE ?", login, email)
	return scanOneUser(row)
}

func GetAllUser() ([]User, error) {
	rows, err := db.Get().Query("SELECT * FROM user ORDER BY login, email ASC")

	if err != nil {
		return nil, err
	}

	return scanUser(rows)
}

func scanUser(rows *sql.Rows) ([]User, error) {
	user := make([]User, 0)

	for rows.Next() {
		u, err := scanOneUser(rows)

		if err != nil {
			return nil, err
		}

		user = append(user, *u)
	}

	return user, nil
}

func scanOneUser(row util.RowScanner) (*User, error) {
	user := User{}

	if err := row.Scan(&user.Id,
		&user.Login,
		&user.Email,
		&user.Pwd); err != nil {
		return nil, err
	}

	return &user, nil
}
