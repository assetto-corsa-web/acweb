package resp

import (
	"net/http/httptest"
	"strings"
	"testing"
)

type testData struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	err := Success(w, 1, "message", testData{"something", 123})

	if err != nil {
		t.Fatal(err)
	}

	respJson := w.Body.String()

	if !strings.Contains(respJson, "something") ||
		!strings.Contains(respJson, "message") ||
		!strings.Contains(respJson, "type\":0") {
		t.Fatal("Response must contain relevant data")
	}
}

func TestFailure(t *testing.T) {
	w := httptest.NewRecorder()
	err := Failure(w, 1, "message", testData{"something", 123})

	if err != nil {
		t.Fatal(err)
	}

	respJson := w.Body.String()

	if !strings.Contains(respJson, "something") ||
		!strings.Contains(respJson, "message") ||
		!strings.Contains(respJson, "type\":1") {
		t.Fatal("Response must contain relevant data")
	}
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	err := Error(w, 1, "message", testData{"something", 123})

	if err != nil {
		t.Fatal(err)
	}

	respJson := w.Body.String()

	if !strings.Contains(respJson, "something") ||
		!strings.Contains(respJson, "message") ||
		!strings.Contains(respJson, "type\":2") {
		t.Fatal("Response must contain relevant data")
	}
}

func TestSendResponseNilData(t *testing.T) {
	w := httptest.NewRecorder()
	err := sendResponse(w, Resp{99, 123, "test", nil, SuccessHttpCode})

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(w.Body.String(), "data\":null") {
		t.Fatal("Data must be null")
	}
}

func TestHttpStatusCodeOK(t *testing.T) {
	w := httptest.NewRecorder()
	err := sendResponse(w, Resp{1, 2, "", nil, SuccessHttpCode})

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != 200 {
		t.Fatal("Response code must be 200")
	}
}

func TestHttpStatusCode500(t *testing.T) {
	SuccessHttpCode = 500
	w := httptest.NewRecorder()
	err := sendResponse(w, Resp{1, 2, "", nil, SuccessHttpCode})

	if err != nil {
		t.Fatal(err)
	}

	if w.Code != 500 {
		t.Fatal("Response code must be 500")
	}

	SuccessHttpCode = 200
}
