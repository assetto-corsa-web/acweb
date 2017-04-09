package model

import (
	"db"
	"github.com/DeKugelschieber/go-util"
)

type Settings struct {
	Id         int64  `json:"id"`
	Folder     string `json:"folder"`
	Executable string `json:"executable"`
	Args       string `json:"args"`
}

func (m *Settings) Save() error {
	if m.Id == 0 {
		res, err := db.Get().Exec("INSERT INTO settings (folder, executable, args) VALUES (?, ?, ?)", m.Folder,
			m.Executable,
			m.Args)

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

	_, err := db.Get().Exec("UPDATE settings SET folder = ?, executable = ?, args = ? WHERE id = ?", m.Folder,
		m.Executable,
		m.Args,
		m.Id)
	return err
}

func GetSettings() (*Settings, error) {
	row := db.Get().QueryRow("SELECT * FROM settings LIMIT 1")
	return scanSettings(row)
}

func scanSettings(row util.RowScanner) (*Settings, error) {
	settings := Settings{}

	if err := row.Scan(&settings.Id,
		&settings.Folder,
		&settings.Executable,
		&settings.Args); err != nil {
		return nil, err
	}

	return &settings, nil
}
