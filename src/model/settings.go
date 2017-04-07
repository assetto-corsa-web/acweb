package model

import (
	"db"
	"github.com/DeKugelschieber/go-util"
)

type Settings struct {
	Id     int64  `json:"id"`
	Folder string `json:"folder"`
	Cmd    string `json:"cmd"`
}

func (m *Settings) Save() error {
	if m.Id == 0 {
		res, err := db.Get().Exec("INSERT INTO settings (folder, command) VALUES (?, ?)", m.Folder,
			m.Cmd)

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

	_, err := db.Get().Exec("UPDATE settings SET folder = ?, command = ? WHERE id = ?", m.Folder,
		m.Cmd,
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
		&settings.Cmd); err != nil {
		return nil, err
	}

	return &settings, nil
}
