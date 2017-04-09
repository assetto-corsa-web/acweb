package instance

import (
	"os/exec"
)

type Instance struct {
	PID           int    `json:"pid"`
	Name          string `json:"name"`
	Configuration int64  `json:"configuration"`
	Cmd           *exec.Cmd
}
