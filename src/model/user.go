package model

import (
	"database/sql"
	"db"
	"github.com/DeKugelschieber/go-util"
)

type User struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pwd   string // do not send this to client
}

func (m *User) Save(tx *sql.Tx) error {
	tx, err := createTxIfRequired(tx)

	if err != nil {
		return err
	}

	if m.Id == 0 {
		res, err := tx.Exec("INSERT INTO user (login, email, password) VALUES (?, ?, ?)", m.Login,
			m.Email,
			m.Pwd)

		if err != nil {
			tx.Rollback()
			return err
		}

		id, err := res.LastInsertId()

		if err != nil {
			tx.Rollback()
			return err
		}

		m.Id = id
		return tx.Commit()
	}

	_, err = tx.Exec("UPDATE user SET login = ?, email = ?, password = ? WHERE id = ?", m.Login,
		m.Email,
		m.Pwd,
		m.Id)
	return tx.Commit()
}

func GetUserByLoginOrEmailAndPassword(login, email, pwd string) (*User, error) {
	row := db.Get().QueryRow("SELECT * FROM user WHERE (login LIKE ? OR email LIKE ?) AND password = ?", login, email, pwd)
	return scanOneUser(row)
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
