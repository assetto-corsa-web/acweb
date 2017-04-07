package settings

import (
	"log"
	"model"
)

func GetSettings() *model.Settings {
	settings, err := model.GetSettings()

	if err != nil {
		log.Printf("Error reading settings: %v", err)
		settings = &model.Settings{}
	}

	return settings
}
