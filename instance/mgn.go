package instance

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Kugelschieber/acweb/model"
	"github.com/Kugelschieber/acweb/settings"
	"github.com/Kugelschieber/acweb/util"
)

var (
	instances = make([]Instance, 0)
	m         sync.Mutex
)

func GetAllInstances() []Instance {
	return instances
}

func StartInstance(instanceName string, configuration int64) error {
	// check input
	instanceName = util.Trim(instanceName)

	if instanceName == "" {
		return util.OpError{1, "Instance name must be set"}
	}

	// create log dir
	if err := os.MkdirAll(os.Getenv("ACWEB_INSTANCE_LOGDIR"), 0755); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error creating server log folder")
		return util.OpError{5, "Error creating server log folder"}
	}

	// read config
	config, err := model.GetConfigurationById(configuration)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading configuration to start instance")
		return util.OpError{2, "Error reading configuration"}
	}

	if err := config.Join(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error joining entities to configuration to start instance")
		return util.OpError{2, "Error reading configuration"}
	}

	// read settings
	s := settings.GetSettings()

	// write config
	iniServerCfg, iniEntryList, err := writeConfig(config)

	if err != nil {
		return util.OpError{3, "Error writing configuration"}
	}

	// force server_cfg and entry_list ini paths
	// start
	// FIXME: s.Args has been  discarted. No real use so far?
	// cmd := exec.Command(filepath.Join(s.Folder, s.Executable), strings.Split(cmdArgs, " ")...)
	cmd := exec.Command(filepath.Join(s.Folder, s.Executable), "-c", iniServerCfg, "-e", iniEntryList)
	now := time.Now().Format("20060102_150405")

	logName := now + "_" + int64ToStr(config.Id) + "_" + instanceName + ".log"
	logfile, err := os.Create(filepath.Join(os.Getenv("ACWEB_INSTANCE_LOGDIR"), logName))

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error creating log file")
		return util.OpError{6, "Error creating log file"}
	}

	cmd.Stdout = logfile
	cmd.Stderr = logfile

	// run acServer from its folder so track and car data will be read for checksum;
	cmd.Dir = s.Folder

	if err := cmd.Start(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error starting instance")
		return util.OpError{4, "Error starting instance"}
	}

	instance := Instance{cmd.Process.Pid, instanceName, config.Id, cmd, logfile}
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
			log.WithFields(log.Fields{"err": err}).Error("Error when instance stopped")
		} else {
			log.WithFields(log.Fields{"err": exitErr.Error(), "process_state": exitErr.ProcessState}).Error("Error when instance stopped")
		}
	}

	// remove process
	m.Lock()
	defer m.Unlock()

	for i, instance := range instances {
		if instance.PID == cmd.Process.Pid {
			if err := instance.File.Close(); err != nil {
				log.WithFields(log.Fields{"err": err}).Error("Error closing log file")
			}

			instances = append(instances[:i], instances[i+1:]...)
			log.WithFields(log.Fields{"instance": i, "pid": instance.PID}).Info("Instance removed")
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
		log.WithFields(log.Fields{"err": err}).Error("Error when stopping instance")
		return util.OpError{1, "Error stopping instance"}
	}

	return nil
}
