package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"golang.org/x/crypto/bcrypt"
)

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("parse form error: %v", err)
	}

	msg := r.FormValue("message")
	real := []byte("$2a$10$2ovnPWuIjMx2S0HvCxP/mutzdsGhyt8rq/JqnJg/6OyC3B0APMGlK")

	if err := bcrypt.CompareHashAndPassword(real, []byte(msg)); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, _ = w.Write([]byte("Nope!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Right!"))
}

func HandleReq(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received: ", r.Method)

	switch r.Method {
	case http.MethodGet:
		display(w, r)
	case http.MethodPost:
		add(w, r)
	default:
		w.WriteHeader(405)
		_, _ = w.Write([]byte("Method not allowed"))
	}
}

func display(w http.ResponseWriter, r *http.Request) {
	log.Println("Response returned: ", 200)
	_, _ = w.Write([]byte("200"))
}

func add(w http.ResponseWriter, r *http.Request) {
	log.Println("Response returned: ", 201)
	_, _ = w.Write([]byte("201"))
}

/*
func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/trace", Trace)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
}
*/

func main() {
	http.HandleFunc("/guess", GuessHandler)
	http.HandleFunc("/puzzle", HandleReq)

	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}
