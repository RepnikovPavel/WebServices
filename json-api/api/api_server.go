package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DoSomeStuff() {
	fmt.Println("some stuff")
}

type ServerConfig struct {
	Sockaddr string
}

type payload func(w http.ResponseWriter, r *http.Request) error

// TODO: make advanced custom api errors
const (
	ApiError0 string = "api error 0"
	ApiError1 string = "api error 1"
)

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.WriteHeader(statusCode)
	// Header signature is
	// type Header map[string][]string
	// Set(key: x string,value: y string)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f payload) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			// TODO: make advanced error handling
			WriteJSON(w, http.StatusBadRequest, ApiError0)
			log.Println(fmt.Sprintf("error: %s\nin %s", err, "/api_server/makeHTTPHadleFunc"))

		}
		// TODO: make advanced custom api errors
	}
}

type Server struct {
}

func (s *Server) Run() {
	router := mux.NewRouter()

	/*
		takes path string and
		handle function
		type HandlerFunc func(ResponseWriter, *Request)
	*/
	router.HandleFunc()
}

// CRUD for accounts
