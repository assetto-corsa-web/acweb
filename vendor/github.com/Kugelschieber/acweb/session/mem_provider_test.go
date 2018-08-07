package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewSession(t *testing.T) {
	testMemCreateValidSession(t)
}

func TestMemCreateCookie(t *testing.T) {
	New("cookieName", 3, NewMemProvider())
	recorder := httptest.NewRecorder()
	_, err := NewSession(recorder, &http.Request{})

	if err != nil {
		t.Fatal(err)
	}

	request := http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	cookie, err := request.Cookie("cookieName")

	if err != nil {
		t.Fatal(err)
	}

	if cookie.Value == "" {
		t.Fatal("Cookie token must be set")
	}
}

func TestMemReadSession(t *testing.T) {
	session := testMemCreateValidSession(t)
	read, err := GetSession(session.Token())

	if err != nil {
		t.Fatal(err)
	}

	if read.Token() != session.Token() {
		t.Fatal("Session tokens must match")
	}
}

func TestMemReadCurrentSession(t *testing.T) {
	New("cookieName", 3, NewMemProvider())
	recorder := httptest.NewRecorder()
	session, err := NewSession(recorder, &http.Request{})

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	read, err := GetCurrentSession(request)

	if err != nil {
		t.Fatal(err)
	}

	if read.Token() != session.Token() {
		t.Fatal("Session tokens must match")
	}
}

func TestMemWriteSession(t *testing.T) {
	session := testMemCreateValidSession(t)
	session.Set("key", "value")
	session.Save()
	read, err := GetSession(session.Token())

	if err != nil {
		t.Fatal(err)
	}

	if read.Token() != session.Token() {
		t.Fatal("Session tokens must match")
	}

	var str string
	read.Get("key", &str)

	if str != "value" {
		t.Fatal("Value does not match saved value")
	}
}

func TestMemDestroySession(t *testing.T) {
	session := testMemCreateValidSession(t)

	session.Destroy(nil, nil)
	_, err := GetSession(session.Token())

	if err == nil {
		t.Fatal("Session must not exist")
	}
}

func TestMemDestroySessionCookie(t *testing.T) {
	New("cookieName", 3, NewMemProvider())
	recorder := httptest.NewRecorder()
	session, err := NewSession(recorder, &http.Request{})

	if !session.Active() {
		t.Fatal("Session token must be set")
	}

	// obtain cookie value
	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	cookie, err := request.Cookie("cookieName")

	if err != nil {
		t.Fatal(err)
	}

	if cookie.Value == "" {
		t.Fatal("Cookie value must be set")
	}

	// destroy and try to read
	err = session.Destroy(recorder, request)

	if err != nil {
		t.Log("Session on destroy not found")
		t.Fatal(err)
	}

	_, err = GetCurrentSession(request)

	if err == nil {
		t.Fatal("Session must not exist")
	}
}

func TestMemGCDestroy(t *testing.T) {
	session := testMemCreateValidSession(t)
	time.Sleep(time.Duration(3050) * time.Millisecond) // 0.05 sec tolerance

	_, err := GetSession(session.Token())

	if err == nil {
		t.Fatal("Session must be destroyed")
	}
}

func TestMemGCExists(t *testing.T) {
	session := testMemCreateValidSession(t)
	time.Sleep(time.Duration(2950) * time.Millisecond) // -0.05 sec tolerance

	_, err := GetSession(session.Token())

	if err != nil {
		t.Fatal(err)
	}
}

func TestMemLifetime(t *testing.T) {
	session := testMemCreateValidSession(t)
	time.Sleep(time.Duration(2900) * time.Millisecond)

	read, err := GetSession(session.Token())

	if err != nil {
		t.Fatal(err)
	}

	if read.lifetime.Before(time.Now().Add(time.Duration(-50) * time.Millisecond)) {
		t.Fatal("Lifetime of session must be updated when read")
	}
}

func testMemCreateValidSession(t *testing.T) Session {
	New("cookieName", 3, NewMemProvider())
	session, err := NewSession(nil, nil)

	if err != nil {
		t.Fatal(err)
	}

	if !session.Active() {
		t.Fatal("Session must be active")
	}

	return session
}
