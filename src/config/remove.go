package config

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
)

func RemoveConfiguration(id int64) error {
	config, err := model.GetConfigurationById(id)

	if err != nil {
		log.Printf("Error reading configuration: %v", err)
		return util.OpError{1, "Error reading configuration"}
	}

	if err := config.Remove(); err != nil {
		log.Printf("Error removing configuration: %v", err)
		return util.OpError{2, "Error removing configuration"}
	}

	return nil
}
