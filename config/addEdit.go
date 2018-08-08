package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/acweb/model"
	"github.com/assetto-corsa-web/acweb/util"
)

func AddEditConfiguration(config *model.Configuration) error {
	// TODO check more fields
	config.Name = util.Trim(config.Name)

	if config.Name == "" {
		return util.OpError{1, "Name must be set"}
	}

	if len(config.Weather) < 1 {
		return util.OpError{2, "There must be at least one weather defined"}
	}

	if len(config.Cars) < 1 {
		return util.OpError{3, "There must be at least one car defined"}
	}

	setDynamicTrack(config)

	// create/edit
	if config.Id != 0 {
		existing, err := model.GetConfigurationById(config.Id)

		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error reading configuration")
			return util.OpError{4, "Error reading configuration"}
		}

		if err := existing.Join(); err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error joining entities to configuration")
			return util.OpError{5, "Error joining entities to configuration"}
		}

		if err := removeWeather(existing, config); err != nil {
			return util.OpError{6, "Error removing weather"}
		}

		if err := removeCars(existing, config); err != nil {
			return util.OpError{7, "Error removing cars"}
		}
	}

	if err := config.Save(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error saving configuration")
		return util.OpError{8, "Error saving configuration"}
	}

	return nil
}

func setDynamicTrack(config *model.Configuration) {
	if !config.DynamicTrack || config.Condition == "CUSTOM" {
		return
	}

	if config.Condition == "DUSTY" {
		config.StartValue = 86
		config.Randomness = 1
		config.TransferredGrip = 50
		config.LapsToImproveGrip = 30
	} else if config.Condition == "OLD" {
		config.StartValue = 89
		config.Randomness = 3
		config.TransferredGrip = 80
		config.LapsToImproveGrip = 50
	} else if config.Condition == "SLOW" {
		config.StartValue = 96
		config.Randomness = 1
		config.TransferredGrip = 80
		config.LapsToImproveGrip = 300
	} else if config.Condition == "GREEN" {
		config.StartValue = 95
		config.Randomness = 2
		config.TransferredGrip = 90
		config.LapsToImproveGrip = 132
	} else if config.Condition == "FAST" {
		config.StartValue = 98
		config.Randomness = 2
		config.TransferredGrip = 80
		config.LapsToImproveGrip = 700
	} else if config.Condition == "OPTIMUM" {
		config.StartValue = 100
		config.Randomness = 0
		config.TransferredGrip = 100
		config.LapsToImproveGrip = 1
	}
}

func removeWeather(old, config *model.Configuration) error {
	for _, oldweather := range old.Weather {
		found := false

		for _, weather := range config.Weather {
			if weather.Id == oldweather.Id {
				found = true
				break
			}
		}

		if !found {
			if err := oldweather.Remove(); err != nil {
				log.WithFields(log.Fields{"err": err}).Error("Error removing weather")
				return err
			}
		}
	}

	return nil
}

func removeCars(old, config *model.Configuration) error {
	for _, oldcar := range old.Cars {
		found := false

		for _, car := range config.Cars {
			if car.Id == oldcar.Id {
				found = true
				break
			}
		}

		if !found {
			if err := oldcar.Remove(); err != nil {
				log.WithFields(log.Fields{"err": err}).Error("Error removing car")
				return err
			}
		}
	}

	return nil
}
