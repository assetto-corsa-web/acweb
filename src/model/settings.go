package model

type Settings struct {
	Id         int64  `json:"id"`
	Folder     string `json:"folder"`
	Executable string `json:"executable"`
	Args       string `json:"args"`
}

func (m *Settings) Save() error {
	if m.Id == 0 {
		res, err := session.NamedExec("INSERT INTO settings (folder, executable, args) VALUES (:folder, :executable, :args)", m)

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

	_, err := session.NamedExec("UPDATE settings SET folder = :folder, executable = :executable, args = :args WHERE id = :id", m)
	return err
}

func GetSettings() (*Settings, error) {
	settings := new(Settings)

	if err := session.Get(settings, "SELECT * FROM settings LIMIT 1"); err != nil {
		return nil, err
	}

	return settings, nil
}
