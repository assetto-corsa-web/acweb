package instance

import (
	"github.com/DeKugelschieber/go-util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetAllInstanceLogs() ([]Log, error) {
	dir, err := ioutil.ReadDir(os.Getenv("ACWEB_INSTANCE_LOGDIR"))

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
	content, err := ioutil.ReadFile(filepath.Join(os.Getenv("ACWEB_INSTANCE_LOGDIR"), file))

	if err != nil {
		log.Printf("Error reading log file: %v", err)
		return "", util.OpError{1, "Error reading log file"}
	}

	return string(content), nil
}
