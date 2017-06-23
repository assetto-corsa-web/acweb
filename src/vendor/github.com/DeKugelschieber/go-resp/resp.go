package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	SUCCESS = 0
	FAILURE = 1
	ERROR   = 2
)

var (
	// When set to true, Failure() and Error() will write to log.
	// This is enabled by default.
	Log = true

	// HTTP codes send when appropriate function is used.
	// All status OK by default.
	SuccessHttpCode = http.StatusOK
	FailureHttpCode = http.StatusOK
	ErrorHttpCode   = http.StatusOK
)

// Response of type "success", "failure" or "error" with
// error code, a short message and additional data.
// The data will be marshalled to JSON and must have decorators
// for all fields which are supposed to be send to the client.
// The type can be used to differentiate between different kind
// of outcomes:
//
// success = everything went fine, the operation was successful
// failure = the operation failed due to wrong input or usage
// error   = a technical error occured
type Resp struct {
	Type     int         `json:"type"`
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	httpCode int
}

// Sends a success response with given code, message and data
// to the client. The first argument must be the ResponseWriter.
func Success(w http.ResponseWriter, code int, msg string, data interface{}) error {
	return sendResponse(w, Resp{SUCCESS, code, msg, data, SuccessHttpCode})
}

// Sends a failure response with given code, message and data
// to the client. The first argument must be the ResponseWriter.
func Failure(w http.ResponseWriter, code int, msg string, data interface{}) error {
	if Log {
		log.Printf("Error %v: %v", code, msg)
	}

	return sendResponse(w, Resp{FAILURE, code, msg, data, FailureHttpCode})
}

// Sends an error response with given code, message and data
// to the client. The first argument must be the ResponseWriter.
func Error(w http.ResponseWriter, code int, msg string, data interface{}) error {
	if Log {
		log.Printf("Error %v: %v", code, msg)
	}

	return sendResponse(w, Resp{ERROR, code, msg, data, ErrorHttpCode})
}

func sendResponse(w http.ResponseWriter, resp Resp) error {
	response, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	w.WriteHeader(resp.httpCode)
	w.Write(response)
	return nil
}
