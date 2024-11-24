package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	got, _ := GetAPIKey(h)
	want := ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
