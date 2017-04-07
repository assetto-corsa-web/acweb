package rest

import (
	"encoding/json"
	"github.com/DeKugelschieber/go-resp"
	"github.com/DeKugelschieber/go-session"
	"github.com/DeKugelschieber/go-util"
	"log"
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
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		log.Print(err)
		resp.Error(w, 100, err.Error(), nil)
		return
	}

	userId, err := user.Login(req.Login, req.Pwd)

	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return
	}

	// start session
	s, err := session.NewSession(w, r)

	if err != nil {
		log.Printf("Error starting session on login: %v", err)
		resp.Error(w, 3, err.Error(), nil)
		return
	}

	s.Set("user_id", userId)

	if err := s.Save(); err != nil {
		log.Printf("Error saving session on login: %v", err)
		resp.Error(w, 4, err.Error(), nil)
		return
	}

	resp.Success(w, 0, "", nil)
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

	resp.Success(w, 0, "", nil)
}

func AddEditUser(w http.ResponseWriter, r *http.Request) {
	var req addEditUserReq
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		log.Print(err)
		resp.Error(w, 100, err.Error(), nil)
		return
	}

	err := user.AddEditUser(req.Id, req.Login, req.Email, req.Pwd1, req.Pwd2)

	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return
	}

	resp.Success(w, 0, "", nil)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	var req removeUserReq
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		log.Print(err)
		resp.Error(w, 100, err.Error(), nil)
		return
	}

	err := user.RemoveUser(req.Id)

	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return
	}

	resp.Success(w, 0, "", nil)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.GetAllUser()

	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return
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

	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return
	}

	resp, _ := json.Marshal(user)
	w.Write(resp)
}

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	var req saveSettingsReq
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		log.Print(err)
		resp.Error(w, 100, err.Error(), nil)
		return
	}

	err := settings.SaveSettings(req.Folder, req.Cmd)

	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return
	}

	resp.Success(w, 0, "", nil)
}

func GetSettings(w http.ResponseWriter, r *http.Request) {
	settings := settings.GetSettings()
	resp, _ := json.Marshal(settings)
	w.Write(resp)
}

func AddEditConfiguration(w http.ResponseWriter, r *http.Request) {

}

func RemoveConfiguration(w http.ResponseWriter, r *http.Request) {

}

func GetAllConfigurations(w http.ResponseWriter, r *http.Request) {

}

func GetConfiguration(w http.ResponseWriter, r *http.Request) {

}

func StartInstance(w http.ResponseWriter, r *http.Request) {

}

func StopInstance(w http.ResponseWriter, r *http.Request) {

}

func GetAllInstances(w http.ResponseWriter, r *http.Request) {

}

func GetInstanceLog(w http.ResponseWriter, r *http.Request) {

}
