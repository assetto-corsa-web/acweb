package rest

import (
	"encoding/json"
	"github.com/DeKugelschieber/go-resp"
	"github.com/DeKugelschieber/go-util"
	"log"
	"net/http"
)

func decode(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		log.Print(err)
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
