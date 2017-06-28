/*
 * Tool to extract track and car information from AC.
 * Ugly but it works.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Track struct {
	Name        string `json:"name"`
	Config      string `json:"config"`
	Description string `json:"description"`
	MaxSlots    int    `json:"max_slots"`
}

type UITrack struct {
	Pitboxes string `json:"pitboxes"`
}

type Car struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Paintings   []string `json:"paintings"`
}

var (
	ac_dir = ""
)

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	ac_dir = flag.Arg(1)

	if ac_dir != "" && cmd == "tracks" {
		extractTracks()
	} else if ac_dir != "" && cmd == "cars" {
		extractCars()
	} else {
		fmt.Println("Usage: ./extractor tracks|cars path_to_assetto_corsa_content_dir")
	}
}

func extractTracks() {
	tracks := make([]Track, 0)
	trackDirs, err := ioutil.ReadDir(filepath.Join(ac_dir, "tracks"))

	if err != nil {
		log.Fatal(err)
	}

	for _, trackDir := range trackDirs {
		// discart non-dir
		if !trackDir.IsDir() {
			continue
		}

		trackName := trackDir.Name()
		trackConfig := ""
		maxSlots := 0
		trackConfigsFound := false

		// read ui dir
		uiDirs, err := ioutil.ReadDir(filepath.Join(ac_dir, "tracks", trackName, "ui"))

		if err != nil {
			log.Fatal(err)
		}

		// subversions of track
		for _, uiDir := range uiDirs {
			if uiDir.IsDir() {
				trackConfig = uiDir.Name()
				maxSlots = getMaxSlots(filepath.Join(ac_dir, "tracks", trackName, "ui", uiDir.Name()))
				trackConfigsFound = true
				tracks = append(tracks, Track{trackName,
					trackConfig,
					title(trackName) + " - " + title(trackConfig),
					maxSlots})
			}
		}

		// no subversions
		if !trackConfigsFound {
			maxSlots = getMaxSlots(filepath.Join(ac_dir, "tracks", trackDir.Name(), "ui"))
			tracks = append(tracks, Track{trackName,
				"",
				title(trackName),
				maxSlots})
		}
	}

	// write to file
	out, err := json.MarshalIndent(tracks, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("tracks.json", out, 0755); err != nil {
		log.Fatal(err)
	}
}

func getMaxSlots(dir string) int {
	var content []byte
	var err error
	path := filepath.Join(dir, "ui_track.json")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		content, err = ioutil.ReadFile(filepath.Join(dir, "dlc_ui_track.json"))
	} else {
		content, err = ioutil.ReadFile(path)
	}

	if err != nil {
		log.Fatal(err)
	}

	uitrack := UITrack{}

	if err := json.Unmarshal(content, &uitrack); err != nil {
		log.Fatal(err)
	}

	maxSlots, err := strconv.Atoi(uitrack.Pitboxes)

	if err != nil {
		log.Fatal(err)
	}

	return maxSlots
}

func title(str string) string {
	if len(str) < 2 {
		return str
	}

	b := []byte(str)
	return strings.ToUpper(string(b[0])) + string(b[1:])
}

func extractCars() {
	cars := make([]Car, 0)
	carDirs, err := ioutil.ReadDir(filepath.Join(ac_dir, "cars"))

	if err != nil {
		log.Fatal(err)
	}

	for _, carDir := range carDirs {
		carName := carDir.Name()
		skins := make([]string, 0)

		// read skins dir
		skinPath := filepath.Join(ac_dir, "cars", carDir.Name(), "skins")

		if _, err := os.Stat(skinPath); err == nil {
			skinsDir, err := ioutil.ReadDir(skinPath)

			if err != nil {
				log.Fatal(err)
			}

			for _, skinDir := range skinsDir {
				skins = append(skins, skinDir.Name())
			}
		}

		// read name for description
		carDesc := getCarName(filepath.Join(ac_dir, "cars", carDir.Name(), "ui"))

		// add
		cars = append(cars, Car{carName, carDesc, skins})
	}

	// write to file
	out, err := json.MarshalIndent(cars, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("cars.json", out, 0755); err != nil {
		log.Fatal(err)
	}
}

func getCarName(dir string) string {
	var content []byte
	var err error
	path := filepath.Join(dir, "ui_car.json")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		content, err = ioutil.ReadFile(filepath.Join(dir, "dlc_ui_car.json"))
	} else {
		content, err = ioutil.ReadFile(path)
	}

	if err != nil {
		log.Fatal(err)
	}

	// ugly, but who cars if the json contains non unicode?
	start := strings.Index(string(content), "\"name\": \"")

	if start < 0 {
		log.Fatal("Start of car name not found in ui_car.json!")
	}

	start += 9
	name := ""

	for {
		name += string(content[start])
		start++

		if content[start] == '"' {
			break
		}
	}

	return name
}
