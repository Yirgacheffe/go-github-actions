package utcase

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	Routes()
}

func TestSendJSON(t *testing.T) {

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotX, err)
		}
		t.Log("\tShould be able to create q request.", checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}

		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}

		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "zeng tai" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "zeng.tai@ardanstudios.com" {
			t.Log("\tShould have a Email.", checkMark)
		} else {
			t.Error("\tShould have a Email.", ballotX, u.Name)
		}
	}
}

func TestSendJSONBasic(t *testing.T) {

	type user struct {
		Name  string
		Email string
	}

	tests := []struct {
		name   string
		param  string
		expect user
	}{
		{"base case", "noparam", user{"zeng tai", "zeng.tai@ardanstudios.com"}},
		{"bad  case", "noparam", user{"zeng tai", "zeng.tai@ardanstudios.com"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/sendjson", nil) // strings.NewReader(tt.param) if you have parameters
			if err != nil {
				t.Fatalf("failed to make request %v", err)
			}

			w := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			u := &user{}
			err = json.NewDecoder(w.Body).Decode(&u) // err := json.Unmarshal([]byte(w.Body.String()), &resp)

			assert.Nil(t, err)
			assert.Equal(t, tt.expect.Name, u.Name)
			assert.Equal(t, tt.expect.Email, u.Email)
		})
	}
}

func TestDeleteEntry(t *testing.T) {
	req, err := http.NewRequest("DELETE", "entries", nil)
	if err != nil {
		t.Fatalf("failed to make request %v", err)
	}

	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(EntryHandler)
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteEntry_405(t *testing.T) {
	req, err := http.NewRequest("POST", "entries", nil)
	if err != nil {
		t.Fatalf("failed to make request %v", err)
	}

	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(EntryHandler)
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}
