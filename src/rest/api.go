package rest

import (
	"config"
	"encoding/json"
	"github.com/DeKugelschieber/go-resp"
	"github.com/DeKugelschieber/go-session"
	"instance"
	"log"
	"model"
	"net/http"
	"settings"
	"strconv"
	"user"
)

func CheckSessionHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := session.GetCurrentSession(r)

	if s.Active() {
		var id int64

		if err := s.Get("user_id", &id); err != nil {
			log.Printf("Error reading user ID: %v", err)
			resp.Error(w, 1, "Error reading user ID", nil)
			return
		}

		resp.Success(w, 0, "", loginRes{id})
	} else {
		// don't log this
		resp.Log = false
		resp.Failure(w, 3, "Not logged in", nil)
		resp.Log = true
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq

	if decode(w, r, &req) {
		return
	}

	user, err := user.Login(req.Login, req.Pwd)

	if iserror(w, err) {
		return
	}

	// start session
	s, err := session.NewSession(w, r)

	if err != nil {
		log.Printf("Error starting session on login: %v", err)
		resp.Error(w, 3, err.Error(), nil)
		return
	}

	s.Set("user_id", user.Id)
	s.Set("admin", user.Admin)
	s.Set("moderator", user.Moderator)

	if err := s.Save(); err != nil {
		log.Printf("Error saving session on login: %v", err)
		resp.Error(w, 4, err.Error(), nil)
		return
	}

	success(w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	s, err := session.GetCurrentSession(r)

	if !s.Active() {
		log.Printf("Session not found on logout: %v", err)
		resp.Error(w, 1, "Session not found", err)
		return
	}

	if err := s.Destroy(w, r); err != nil {
		log.Printf("Error destroying user session on logout: %v", err)
		resp.Error(w, 2, "Error destroying user session", nil)
		return
	}

	success(w)
}

func AddEditUser(w http.ResponseWriter, r *http.Request) {
	if !isadmin(r) {
		denyAccess(w)
		return
	}

	var req addEditUserReq

	if decode(w, r, &req) {
		return
	}

	err := user.AddEditUser(req.Id,
		req.Login,
		req.Email,
		req.Pwd1,
		req.Pwd2,
		req.Admin,
		req.Moderator)

	if iserror(w, err) {
		return
	}

	success(w)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	if !isadmin(r) {
		denyAccess(w)
		return
	}

	var req removeUserReq

	if decode(w, r, &req) {
		return
	}

	err := user.RemoveUser(req.Id)

	if iserror(w, err) {
		return
	}

	success(w)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.GetAllUser()

	if iserror(w, err) {
		return
	}

	if !isadmin(r) {
		for i := range user {
			user[i].Email = ""
		}
	}

	resp, _ := json.Marshal(user)
	w.Write(resp)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		resp.Error(w, 100, err.Error(), nil)
		return
	}

	user, err := user.GetUser(int64(id))

	if iserror(w, err) {
		return
	}

	if !isadmin(r) {
		user.Email = ""
	}

	resp, _ := json.Marshal(user)
	w.Write(resp)
}

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	if !isadmin(r) {
		denyAccess(w)
		return
	}

	var req saveSettingsReq

	if decode(w, r, &req) {
		return
	}

	err := settings.SaveSettings(req.Folder, req.Executable, req.Args)

	if iserror(w, err) {
		return
	}

	success(w)
}

func GetSettings(w http.ResponseWriter, r *http.Request) {
	settings := settings.GetSettings()
	resp, _ := json.Marshal(settings)
	w.Write(resp)
}

func AddEditConfiguration(w http.ResponseWriter, r *http.Request) {
	if !ismoderator(r) {
		denyAccess(w)
		return
	}

	var req model.Configuration

	if decode(w, r, &req) {
		return
	}

	err := config.AddEditConfiguration(&req)

	if iserror(w, err) {
		return
	}

	success(w)
}

func RemoveConfiguration(w http.ResponseWriter, r *http.Request) {
	if !ismoderator(r) {
		denyAccess(w)
		return
	}

	var req removeConfigReq

	if decode(w, r, &req) {
		return
	}

	err := config.RemoveConfiguration(req.Id)

	if iserror(w, err) {
		return
	}

	success(w)
}

func GetAllConfigurations(w http.ResponseWriter, r *http.Request) {
	configs, err := config.GetAllConfigurations()

	if iserror(w, err) {
		return
	}

	resp, _ := json.Marshal(configs)
	w.Write(resp)
}

func GetConfiguration(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		resp.Error(w, 100, err.Error(), nil)
		return
	}

	config, err := config.GetConfiguration(int64(id))

	if iserror(w, err) {
		return
	}

	resp, _ := json.Marshal(config)
	w.Write(resp)
}

func GetAvailableTracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := config.GetAvailableTracks()

	if iserror(w, err) {
		return
	}

	resp, _ := json.Marshal(tracks)
	w.Write(resp)
}

func GetAvailableCars(w http.ResponseWriter, r *http.Request) {
	cars, err := config.GetAvailableCars()

	if iserror(w, err) {
		return
	}

	resp, _ := json.Marshal(cars)
	w.Write(resp)
}

func StartInstance(w http.ResponseWriter, r *http.Request) {
	if !ismoderator(r) {
		denyAccess(w)
		return
	}

	var req startInstanceReq

	if decode(w, r, &req) {
		return
	}

	err := instance.StartInstance(req.Name, req.Configuration)

	if iserror(w, err) {
		return
	}

	success(w)
}

func StopInstance(w http.ResponseWriter, r *http.Request) {
	if !ismoderator(r) {
		denyAccess(w)
		return
	}

	var req stopInstanceReq

	if decode(w, r, &req) {
		return
	}

	err := instance.StopInstance(req.PID)

	if iserror(w, err) {
		return
	}

	success(w)
}

func GetAllInstances(w http.ResponseWriter, r *http.Request) {
	instances := instance.GetAllInstances()
	resp, _ := json.Marshal(instances)
	w.Write(resp)
}

func GetAllInstanceLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := instance.GetAllInstanceLogs()

	if iserror(w, err) {
		return
	}

	resp, _ := json.Marshal(logs)
	w.Write(resp)
}

func GetInstanceLog(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	log, err := instance.GetInstanceLog(file)

	if iserror(w, err) {
		return
	}

	resp, _ := json.Marshal(log)
	w.Write(resp)
}
