package instance

import (
	"github.com/DeKugelschieber/go-util"
	"log"
	"model"
	"os/exec"
	"path/filepath"
	"settings"
	"strings"
	"sync"
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

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting instance: %v", err)
		return util.OpError{4, "Error starting instance"}
	}

	instance := Instance{cmd.Process.Pid, name, config.Id, cmd}
	m.Lock()
	instances = append(instances, instance)
	m.Unlock()
	go observeProcess(cmd)

	return nil
}

func observeProcess(cmd *exec.Cmd) {
	if err := cmd.Wait(); err != nil {
		log.Printf("Error when instance stopped: %v", err)
	}

	// remove process
	m.Lock()
	defer m.Unlock()

	for i, instance := range instances {
		if instance.PID == cmd.Process.Pid {
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
