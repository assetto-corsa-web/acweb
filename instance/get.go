package instance

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/assetto-corsa-web/acweb/util"
)

const (
	max_log_size = 256000 // 256kb
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

func GetInstanceLog(filename string) (string, error) {
	file, err := os.Open(filepath.Join(os.Getenv("ACWEB_INSTANCE_LOGDIR"), filename))

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error opening log file")
		return "", util.OpError{3, "Error opening log file"}
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error obtaining log file info")
		return "", util.OpError{4, "Error obtaining log file info"}
	}

	start := info.Size() - max_log_size + 1 // plus one to read last character

	if start < 0 {
		start = 0
	}

	content := make([]byte, max_log_size)
	_, err = file.ReadAt(content, start)

	if err != nil && err != io.EOF {
		log.WithFields(log.Fields{"err": err}).Error("Error reading log file")
		return "", util.OpError{5, "Error reading log file"}
	}

	return string(content), nil
}
