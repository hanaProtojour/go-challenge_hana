package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TesthandlePost(t *testing.T) {
	expected := `{"seeds" : ["abc", "def", "xyz"]}`
	req := httptest.NewRequest(http.MethodPost, `{"seeds" : ["abc", "def", "xyz"]}`, nil)
	w := httptest.NewRecorder()
	HandlePost(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("oh no", string(data))
	}
}

func TestErrorJSONPost(t *testing.T) {
	
}
