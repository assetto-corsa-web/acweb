package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/acweb/model"
	"github.com/assetto-corsa-web/acweb/util"
)

const (
	track_file = "tracks.json"
	car_file   = "cars.json"
)

func GetAllConfigurations() ([]model.Configuration, error) {
	configs, err := model.GetAllConfigurations()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading all configurations")
		return nil, util.OpError{1, "Error reading all configurations"}
	}

	return configs, nil
}

func GetConfiguration(id int64) (*model.Configuration, error) {
	config, err := model.GetConfigurationById(id)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading configuration")
		return nil, util.OpError{1, "Error reading configuration"}
	}

	if err := config.Join(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error joining entities for configuration")
		return nil, util.OpError{2, "Error joining entities for configuration"}
	}

	return config, nil
}

func GetAvailableTracks() ([]Track, error) {
	var tracks []Track
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("ACWEB_CONFIG_DIR"), track_file))

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading track file")
		return nil, util.OpError{1, "Error reading track file"}
	}

	if err := json.Unmarshal(content, &tracks); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error parsing track file")
		return nil, util.OpError{2, "Error parsing track file"}
	}

	return tracks, nil
}

func GetAvailableCars() ([]Car, error) {
	var cars []Car
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("ACWEB_CONFIG_DIR"), car_file))

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading car file")
		return nil, util.OpError{1, "Error reading car file"}
	}

	if err := json.Unmarshal(content, &cars); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error parsing car file")
		return nil, util.OpError{2, "Error parsing car file"}
	}

	return cars, nil
}
