package instance

import (
	"archive/zip"
	"io/ioutil"
	"model"
	"net/http"
	"os"
	"path/filepath"
	"time"

	util "github.com/DeKugelschieber/go-util"
	log "github.com/sirupsen/logrus"
)

func addDataToZip(zw *zip.Writer, filename string, dat []byte) bool {
	fh := &zip.FileHeader{
		Name:   filename,
		Method: zip.Deflate,
	}
	fh.SetModTime(time.Now())

	fw, err := zw.CreateHeader(fh)
	if err != nil {
		log.Printf("Error creating zip header: %v", err)
		return false
	}

	_, err = fw.Write(dat)
	if err != nil {
		log.Printf("Error writing %s config to zip: %v", filename, err)
		return false
	}

	return true
}

func addFileToZip(zw *zip.Writer, filename string, iniFilePath string) bool {
	dat, err := ioutil.ReadFile(iniFilePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return false
	}
	return addDataToZip(zw, filename, dat)
}

// ZipConfiguration creates zip stream with config into and write it into a http.ResponseWriter
func ZipConfiguration(config *model.Configuration, w http.ResponseWriter) error {
	zw := zip.NewWriter(w)
	defer zw.Close()

	serverCfg := ServerConfigToIniString(config)
	if !addDataToZip(zw, ServerIni, []byte(serverCfg)) {
		return util.OpError{1, "Error writing server_cfg.ini into zip archive"}
	}

	entryList := EntryListToIniString(config)
	if !addDataToZip(zw, EntryListIni, []byte(entryList)) {
		return util.OpError{1, "Error writing entry_list.ini into zip archive"}
	}

	return nil
}

// ZipInstanceFiles creates zip stream with current instance and write it into a http.ResponseWriter
func ZipInstanceFiles(config *model.Configuration, w http.ResponseWriter) error {
	zw := zip.NewWriter(w)
	defer zw.Close()

	iniServerCfg := GetServerCfgPath(config)
	if !addFileToZip(zw, ServerIni, iniServerCfg) {
		return util.OpError{1, "Error writing server config into zip archive"}
	}

	iniEntryList := GetEntryListPath(config)
	if !addFileToZip(zw, EntryListIni, iniEntryList) {
		return util.OpError{1, "Error writing entry list into zip archive"}
	}

	return nil
}

// ZipLogFile creates zip stream with a given log file and write it into a http.ResponseWriter
func ZipLogFile(fileName string, w http.ResponseWriter) error {
	zw := zip.NewWriter(w)
	defer zw.Close()

	logFullPath := filepath.Join(os.Getenv("ACWEB_INSTANCE_LOGDIR"), fileName)
	if !addFileToZip(zw, logFullPath, fileName) {
		return util.OpError{1, "Error writing log file into zip archive"}
	}

	return nil
}
