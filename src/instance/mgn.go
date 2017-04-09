package instance

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
	"os"
	"os/exec"
	"path/filepath"
	"settings"
	"strings"
	"sync"
	"time"
)

const (
	log_dir = "server_log"
)

var (
	instances = make([]Instance, 0)
	m         sync.Mutex
)

func GetAllInstances() []Instance {
	return instances
}

func StartInstance(name string, configuration int64) error {
	// check input
	name = util.Trim(name)

	if name == "" {
		return util.OpError{1, "Name must be set"}
	}

	// create log dir
	if err := os.MkdirAll(log_dir, 0755); err != nil {
		log.Printf("Error creating server log folder: %v", err)
		return util.OpError{5, "Error creating server log folder"}
	}

	// read config
	config, err := model.GetConfigurationById(configuration)

	if err != nil {
		log.Printf("Error reading configuration to start instance: %v", err)
		return util.OpError{2, "Error reading configuration"}
	}

	if err := config.Join(); err != nil {
		log.Printf("Error joining entities to configuration to start instance: %v", err)
		return util.OpError{2, "Error reading configuration"}
	}

	// read settings
	s := settings.GetSettings()

	// write config
	if err := writeConfig(s, config); err != nil {
		return util.OpError{3, "Error writing configuration"}
	}

	// start
	cmd := exec.Command(filepath.Join(s.Folder, s.Executable), strings.Split(s.Args, " ")...)
	now := strings.Replace(time.Now().String(), " ", "_", -1)

	logfile, err := os.Create(filepath.Join(log_dir, config.Name+now))

	if err != nil {
		log.Printf("Error creating log file: %v", err)
		return util.OpError{6, "Error creating log file"}
	}

	cmd.Stdout = logfile
	cmd.Stderr = logfile

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting instance: %v", err)
		return util.OpError{4, "Error starting instance"}
	}

	instance := Instance{cmd.Process.Pid, name, config.Id, cmd, logfile}
	m.Lock()
	instances = append(instances, instance)
	m.Unlock()
	go observeProcess(cmd)

	return nil
}

func observeProcess(cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		exitErr, ok := err.(*exec.ExitError)

		if !ok {
			log.Printf("Error when instance stopped: %v", err)
		} else {
			log.Printf("Error when instance stopped: %v %v %v", exitErr.Error(), exitErr.ProcessState, string(exitErr.Stderr))
		}
	}

	// remove process
	m.Lock()
	defer m.Unlock()

	for i, instance := range instances {
		if instance.PID == cmd.Process.Pid {
			if err := instance.File.Close(); err != nil {
				log.Printf("Error closing log file: %v", err)
			}

			instances = append(instances[:i], instances[i+1:]...)
			log.Printf("Instance %v with PID %v removed", i, instance.PID)
			return
		}
	}
}

func StopInstance(pid int) error {
	m.Lock()
	defer m.Unlock()

	for _, instance := range instances {
		if instance.PID == pid {
			// instance is removed from instances by observeProcess
			return stopProcess(&instance)
		}
	}

	return util.OpError{2, "Instance not found"}
}

func stopProcess(instance *Instance) error {
	// just kill it
	if err := instance.Cmd.Process.Kill(); err != nil {
		log.Printf("Error when stopping instance: %v", err)
		return util.OpError{1, "Error stopping instance"}
	}

	return nil
}
