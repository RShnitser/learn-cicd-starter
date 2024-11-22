package auth

import (
    "testing"
	"reflect"
	"net/http"
)

func TestGetAPIKey(t *testing.T) { 
	h := http.Header{}
    got, _ := GetAPIKey(h)
    want := ""
    if !reflect.DeepEqual(want, got) {
         t.Fatalf("expected: %v, got: %v", want, got)
    }
}