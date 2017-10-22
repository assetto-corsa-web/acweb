package model

const (
	mysql_settings_save   = "INSERT INTO settings (folder, executable, args) VALUES (:folder, :executable, :args)"
	mysql_settings_update = "UPDATE settings SET folder = :folder, executable = :executable, args = :args WHERE id = :id"
	mysql_settings_get    = "SELECT * FROM settings LIMIT 1"

	postgres_settings_save   = "INSERT INTO \"settings\" (folder, executable, args) VALUES (:folder, :executable, :args) RETURNING id"
	postgres_settings_update = "UPDATE \"settings\" SET folder = :folder, executable = :executable, args = :args WHERE id = :id"
	postgres_settings_get    = "SELECT * FROM \"settings\" LIMIT 1"
)

type Settings struct {
	Id         int64  `json:"id"`
	Folder     string `json:"folder"`
	Executable string `json:"executable"`
	Args       string `json:"args"`
}

func (s *Settings) Save() error {
	if s.Id == 0 {
		if GetDBType() == "mysql" {
			res, err := session.NamedExec(mysql_settings_save, s)

			if err != nil {
				return err
			}

			id, err := res.LastInsertId()

			if err != nil {
				return err
			}

			s.Id = id
		} else {
			rows, err := session.NamedQuery(postgres_settings_save, s)

			if err != nil {
				return err
			}

			if rows.Next() {
				rows.Scan(&s.Id)
			}

			rows.Close()
		}

		return nil
	}

	var err error

	if GetDBType() == "mysql" {
		_, err = session.NamedExec(mysql_settings_update, s)
	} else {
		_, err = session.NamedExec(postgres_settings_update, s)
	}

	return err
}

func GetSettings() (*Settings, error) {
	settings := new(Settings)

	if GetDBType() == "mysql" {
		if err := session.Get(settings, mysql_settings_get); err != nil {
			return nil, err
		}
	} else {
		if err := session.Get(settings, postgres_settings_get); err != nil {
			return nil, err
		}
	}

	return settings, nil
}
