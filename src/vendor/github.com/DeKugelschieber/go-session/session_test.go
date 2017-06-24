package session

import (
	"testing"
)

type testTypeDeep struct {
	This         string
	Is           float32
	An           int
	EmbeddedType string
	unexported   int
}

type testType struct {
	I          string
	Am         int
	Complex    testTypeDeep
	unexported int
}

func TestSessionDataSetAndGet(t *testing.T) {
	s := Session{}
	s.data = make(map[string]interface{})
	value1 := float64(1.23)
	value2 := "value2"
	value3 := testType{"str", 123, testTypeDeep{"1", 2, 3, "4", 321}, 321}

	s.Set("v1", value1)
	s.Set("v2", value2)
	s.Set("v3", &value3)

	var received1 float64

	if err := s.Get("v1", &received1); err != nil {
		t.Fatalf("Error getting v1: %v", err)
	}

	if received1 != value1 {
		t.Fatalf("received1 was %v", received1)
	}

	var received2 string

	if err := s.Get("v2", &received2); err != nil {
		t.Fatalf("Error getting v2: %v", err)
	}

	if received2 != value2 {
		t.Fatalf("received2 was %v", received2)
	}

	var received3 *testType // FIXME this is ugly

	if err := s.Get("v3", &received3); err != nil {
		t.Fatalf("Error getting v3: %v", err)
	}
}

func TestSessionDataReplaceAndGet(t *testing.T) {
	s := Session{}
	s.data = make(map[string]interface{})

	s.Set("key", "value")
	s.Set("key", "other_value")

	var received string

	if err := s.Get("key", &received); err != nil {
		t.Fatalf("Error getting value: %v", err)
	}

	if received != "other_value" {
		t.Fatalf("received was %v", received)
	}
}

func TestSessionDataRemove(t *testing.T) {
	s := Session{}
	s.data = make(map[string]interface{})

	s.Set("key", "value")
	s.Remove("key")

	var received string

	if err := s.Get("key", &received); err.Error() != "Value not found" {
		t.Fatal("value must not be found")
	}
}

func TestSessionDataNoPanic(t *testing.T) {
	s := Session{}
	s.data = make(map[string]interface{})

	// nil
	if err := s.Get("key", nil); err.Error() != "Value must not be nil" {
		t.Fatal("Value must not be nil")
	}

	// value not found
	if err := s.Get("key", ""); err.Error() != "Value not found" {
		t.Fatal("Value not found")
	}

	// not a pointer
	s.Set("key", "value")

	if err := s.Get("key", ""); err.Error() != "Value type invalid (see https://golang.org/pkg/reflect/#Value.Elem)" {
		t.Fatal("Value type invalid")
	}
}
