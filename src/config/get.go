package config

import (
	"encoding/json"
	"github.com/DeKugelschieber/go-util"
	"io/ioutil"
	"log"
	"model"
	"os"
	"path/filepath"
)

const (
	track_file = "tracks.json"
	car_file   = "cars.json"
)

func GetAllConfigurations() ([]model.Configuration, error) {
	configs, err := model.GetAllConfigurations()

	if err != nil {
		log.Printf("Error reading all configurations: %v", err)
		return nil, util.OpError{1, "Error reading all configurations"}
	}

	return configs, nil
}

func GetConfiguration(id int64) (*model.Configuration, error) {
	config, err := model.GetConfigurationById(id)

	if err != nil {
		log.Printf("Error reading configuration: %v", err)
		return nil, util.OpError{1, "Error reading configuration"}
	}

	if err := config.Join(); err != nil {
		log.Printf("Could not join entities for configuration: %v", err)
		return nil, util.OpError{2, "Error joining entities for configuration"}
	}

	return config, nil
}

func GetAvailableTracks() ([]Track, error) {
	var tracks []Track
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("ACWEB_CONFIG_DIR"), track_file))

	if err != nil {
		log.Printf("Error reading track file: %v", err)
		return nil, util.OpError{1, "Error reading track file"}
	}

	if err := json.Unmarshal(content, &tracks); err != nil {
		log.Printf("Error parsing track file: %v", err)
		return nil, util.OpError{2, "Error parsing track file"}
	}

	return tracks, nil
}

func GetAvailableCars() ([]Car, error) {
	var cars []Car
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("ACWEB_CONFIG_DIR"), car_file))

	if err != nil {
		log.Printf("Error reading car file: %v", err)
		return nil, util.OpError{1, "Error reading car file"}
	}

	if err := json.Unmarshal(content, &cars); err != nil {
		log.Printf("Error parsing car file: %v", err)
		return nil, util.OpError{2, "Error parsing car file"}
	}

	return cars, nil
}
