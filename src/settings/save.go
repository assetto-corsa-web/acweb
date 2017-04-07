package settings

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
)

func SaveSettings(folder, cmd string) error {
	folder = util.Trim(folder)
	cmd = util.Trim(cmd)

	if folder == "" || cmd == "" {
		return util.OpError{1, "Folder and command must be set"}
	}

	settings, err := model.GetSettings()

	if err != nil {
		log.Printf("Error reading settings: %v", err)
		settings = &model.Settings{}
	}

	settings.Folder = folder
	settings.Cmd = cmd

	if err := settings.Save(); err != nil {
		log.Printf("Error saving settings: %v", err)
		return util.OpError{2, "Error saving settings"}
	}

	return nil
}
