package instance

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/acweb/util"
)

func GetAllInstanceLogs() ([]Log, error) {
	dir, err := ioutil.ReadDir(os.Getenv("ACWEB_INSTANCE_LOGDIR"))

	// try to create log directory if it does not exist
	if os.IsNotExist(err) {
		if err := os.MkdirAll(os.Getenv("ACWEB_INSTANCE_LOGDIR"), 0666); err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error creating log directory")
			return nil, util.OpError{1, "Error creating log directory"}
		}

		err = nil
	}

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading log directory")
		return nil, util.OpError{2, "Error reading log directory"}
	}

	logs := make([]Log, 0)

	for _, file := range dir {
		log := Log{file.Name(), file.ModTime(), file.Size()}
		logs = append(logs, log)
	}

	return logs, nil
}

func GetInstanceLog(file string) (string, error) {
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("ACWEB_INSTANCE_LOGDIR"), file))

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading log file")
		return "", util.OpError{3, "Error reading log file"}
	}

	return string(content), nil
}
