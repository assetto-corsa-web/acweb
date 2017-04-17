package instance

import (
	"github.com/DeKugelschieber/go-util"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetAllInstanceLogs() ([]Log, error) {
	dir, err := ioutil.ReadDir(log_dir)

	if err != nil {
		log.Printf("Error reading log directory", err)
		return nil, util.OpError{1, "Error reading log directory"}
	}

	logs := make([]Log, 0)

	for _, file := range dir {
		log := Log{file.Name(), file.ModTime(), file.Size()}
		logs = append(logs, log)
	}

	return logs, nil
}

func GetInstanceLog(file string) (string, error) {
	content, err := ioutil.ReadFile(filepath.Join(log_dir, file))

	if err != nil {
		log.Printf("Error reading log file: %v", err)
		return "", util.OpError{1, "Error reading log file"}
	}

	return string(content), nil
}
