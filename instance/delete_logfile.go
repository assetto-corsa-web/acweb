package instance

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/acweb/util"
)

func DeleteLogFile(filename string) error {
	if err := os.Remove(filepath.Join(os.Getenv("ACWEB_INSTANCE_LOGDIR"), filename)); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error deleting log file")
		return util.OpError{1, "Error deleting log file"}
	}

	return nil
}

func DeleteAllLogFiles() error {
	dir, err := ioutil.ReadDir(os.Getenv("ACWEB_INSTANCE_LOGDIR"))

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error opening log directory")
		return util.OpError{1, "Error opening log directory"}
	}

	for _, file := range dir {
		if !file.IsDir() {
			if err := DeleteLogFile(file.Name()); err != nil {
				log.WithFields(log.Fields{"err": err}).Warn("Error deleting log file when deleting all log files")
			}
		}
	}

	return nil
}
