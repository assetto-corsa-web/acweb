package instance

import (
	"os"
	"os/exec"
)

type Instance struct {
	PID           int    `json:"pid"`
	Name          string `json:"name"`
	Configuration int64  `json:"configuration"`
	Cmd           *exec.Cmd
	File          *os.File
}
