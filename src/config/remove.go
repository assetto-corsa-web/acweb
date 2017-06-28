package config

import (
	"github.com/DeKugelschieber/go-util"
	log "github.com/sirupsen/logrus"
	"model"
)

func RemoveConfiguration(id int64) error {
	config, err := model.GetConfigurationById(id)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading configuration")
		return util.OpError{1, "Error reading configuration"}
	}

	if err := config.Remove(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error removing configuration")
		return util.OpError{2, "Error removing configuration"}
	}

	return nil
}
