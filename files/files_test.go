package files

import (
	"reflect"
	"testing"
)

var (
	want_api_id = "lkajekl"
	want_hits   = "5"
)

func TestReader(t *testing.T) {
	got := Reader()
	want := reflect.TypeOf([]Data{})

	if reflect.TypeOf(got) != want && (len(got) < 1) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestParseLine(t *testing.T) {
	line := `{"api_id":"lkaJek1","hits":5}`
	api_id, hits := ParseLine(line)

	if !(api_id == want_api_id && hits == want_hits) {
		t.Errorf("Expected (%s, %s), but got (%s, %s)", want_api_id, want_hits, api_id, hits)
	}
}

func TestMakeDataCopy(t *testing.T) {
	got := MakeDataCopy(want_api_id, want_hits)
	want := reflect.TypeOf([]Data{})

	if reflect.TypeOf(got) != want && (len(got) < 1) {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
