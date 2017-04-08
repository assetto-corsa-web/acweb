package config

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
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

	// create/edit
	if config.Id != 0 {
		existing, err := model.GetConfigurationById(config.Id)

		if err != nil {
			log.Printf("Error reading configuration: %v", err)
			return util.OpError{4, "Error reading configuration"}
		}

		if err := existing.Join(); err != nil {
			log.Printf("Error joining entities to configuration: %v", err)
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
		log.Printf("Error saving configuration: %v", err)
		return util.OpError{8, "Error saving configuration"}
	}

	return nil
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
				log.Printf("Error removing weather: %v", err)
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
				log.Printf("Error removing car: %v", err)
				return err
			}
		}
	}

	return nil
}
