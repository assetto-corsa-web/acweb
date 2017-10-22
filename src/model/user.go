package model

import (
	"database/sql"
	"errors"
)

const (
	mysql_user_save                = "INSERT INTO user (login, email, password, admin, moderator) VALUES (:login, :email, :password, :admin, :moderator)"
	mysql_user_update              = "UPDATE user SET login = :login, email = :email, password = :password, admin = :admin, moderator = :moderator WHERE id = :id"
	mysql_user_delete              = "DELETE FROM user WHERE id = :id"
	mysql_user_get_login_email_pwd = "SELECT * FROM user WHERE (login LIKE ? OR email LIKE ?) AND password = ?"
	mysql_user_get_id              = "SELECT * FROM user WHERE id = ?"
	mysql_find_user_login_email    = "SELECT * FROM user WHERE login LIKE ? OR email LIKE ?"
	mysql_user_get_all             = "SELECT * FROM user ORDER BY login, email ASC"

	postgres_user_save                = "INSERT INTO \"user\" (login, email, password, admin, moderator) VALUES (:login, :email, :password, :admin, :moderator) RETURNING id"
	postgres_user_update              = "UPDATE \"user\" SET login = :login, email = :email, password = :password, admin = :admin, moderator = :moderator WHERE id = :id"
	postgres_user_delete              = "DELETE FROM \"user\" WHERE id = :id"
	postgres_user_get_login_email_pwd = "SELECT * FROM \"user\" WHERE (login LIKE $1 OR email LIKE $2) AND password = $3"
	postgres_user_get_id              = "SELECT * FROM \"user\" WHERE id = $1"
	postgres_find_user_login_email    = "SELECT * FROM \"user\" WHERE login LIKE $1 OR email LIKE $1"
	postgres_user_get_all             = "SELECT * FROM \"user\" ORDER BY login, email ASC"
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
		var res sql.Result
		var err error

		if GetDBType() == "mysql" {
			res, err = session.NamedExec(mysql_user_save, m)

			if err != nil {
				return err
			}

			id, err := res.LastInsertId()

			if err != nil {
				return err
			}

			m.Id = id
		} else {
			rows, err := session.NamedQuery(postgres_user_save, m)

			if err != nil {
				return err
			}

			if rows.Next() {
				rows.Scan(&m.Id)
			}

			rows.Close()
		}

		return nil
	}

	var err error

	if GetDBType() == "mysql" {
		_, err = session.NamedExec(mysql_user_update, m)
	} else {
		_, err = session.NamedExec(postgres_user_update, m)
	}

	return err
}

func (m *User) Remove() error {
	if m.Id == 0 {
		return errors.New("ID must be set")
	}

	var err error

	if GetDBType() == "mysql" {
		_, err = session.NamedExec(mysql_user_delete, m)
	} else {
		_, err = session.NamedExec(postgres_user_delete, m)
	}

	return err
}

func GetUserByLoginOrEmailAndPassword(login, email, pwd string) (*User, error) {
	user := new(User)

	if GetDBType() == "mysql" {
		if err := session.Get(user, mysql_user_get_login_email_pwd, login, email, pwd); err != nil {
			return nil, err
		}
	} else {
		if err := session.Get(user, postgres_user_get_login_email_pwd, login, email, pwd); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func GetUserById(id int64) (*User, error) {
	user := new(User)

	if GetDBType() == "mysql" {
		if err := session.Get(user, mysql_user_get_id, id); err != nil {
			return nil, err
		}
	} else {
		if err := session.Get(user, postgres_user_get_id, id); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func FindUserByLoginOrEmail(login, email string) (*User, error) {
	user := new(User)

	if GetDBType() == "mysql" {
		if err := session.Get(user, mysql_find_user_login_email, login, email); err != nil {
			return nil, err
		}
	} else {
		if err := session.Get(user, postgres_find_user_login_email, login, email); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func GetAllUser() ([]User, error) {
	user := make([]User, 0)

	if GetDBType() == "mysql" {
		if err := session.Select(&user, mysql_user_get_all); err != nil {
			return nil, err
		}
	} else {
		if err := session.Select(&user, postgres_user_get_all); err != nil {
			return nil, err
		}
	}

	return user, nil
}
