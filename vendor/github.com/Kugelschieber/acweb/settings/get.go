package settings

import (
	log "github.com/sirupsen/logrus"

	"github.com/Kugelschieber/acweb/model"
)

func GetSettings() *model.Settings {
	settings, err := model.GetSettings()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading settings")
		settings = &model.Settings{}
	}

	return settings
}
