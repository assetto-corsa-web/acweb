package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Kugelschieber/acweb/resp"
	"github.com/Kugelschieber/acweb/session"
	"github.com/Kugelschieber/acweb/util"
)

func decode(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error decoding request")
		resp.Error(w, 100, err.Error(), nil)
		return true
	}

	return false
}

func success(w http.ResponseWriter) {
	resp.Success(w, 0, "", nil)
}

func iserror(w http.ResponseWriter, err error) bool {
	if err != nil {
		operr, _ := err.(util.OpError)
		resp.Error(w, operr.Code, operr.Msg, nil)
		return true
	}

	return false
}

func isadmin(r *http.Request) bool {
	s, _ := session.GetCurrentSession(r)
	var admin bool
	s.Get("admin", &admin)

	return admin
}

func ismoderator(r *http.Request) bool {
	s, _ := session.GetCurrentSession(r)
	var admin, moderator bool
	s.Get("admin", &admin)

	if admin {
		return true
	}

	s.Get("moderator", &moderator)

	return moderator
}

func denyAccess(w http.ResponseWriter) {
	resp.Error(w, 200, "Access denied", nil)
}
