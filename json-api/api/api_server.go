package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// api errors segment
type ApiError struct {
	err string
}

func (ApiErr ApiError) Error() string {
	return ApiErr.err
}

const (
	ApiError0                        string = "api error 0"
	ApiError1                        string = "api error 1"
	ApiErrorUnknownHTTPRequestMethod string = "api error UnknownHTTPRequestMethod"
)

// end of epi errors segment

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.WriteHeader(statusCode)
	// Header signature is
	// type Header map[string][]string
	// Set(key: x string,value: y string)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type payload func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f payload) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload_err := f(w, r)
		if payload_err != nil {
			// TODO: make advanced error handling
			json_err := WriteJSON(w, http.StatusBadRequest, ApiError0)
			log.Printf("error: %s\nin %s", payload_err, "/api_server/makeHTTPHadleFunc/payload_func")
			if json_err != nil {
				log.Printf("error: %s\nin %s", json_err, "/api_server/makeHTTPHadleFunc/payload_function/WriteJSON_func")
			}
		}
	}
}

type ServerConfig struct {
	Sockaddr string
}

type Server struct {
	Sockaddr string
}

func NewServer(conf ServerConfig) *Server {
	return &Server{
		Sockaddr: conf.Sockaddr,
	}
}

const (
	rout0 string = "rout0"
	rout1 string = "rout1"
	rout2 string = "rout2"
	rout3 string = "rout3"
	rout4 string = "rout4"
)

// CRUD for accounts
// functions signature is payload function:
// type payload func(w http.ResponseWriter, r *http.Request) error

func PostAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func PutAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HandleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return PostAccount(w, r)
	case "DELETE":
		return DeleteAccount(w, r)
	case "PUT":
		return PutAccount(w, r)
	case "GET":
		return GetAccount(w, r)
	default:
		defer log.Printf("warning: caller side error in %s\ninput HTTP method %s is not supported", "HandleAccount", r.Method)
		return ApiError{err: ApiErrorUnknownHTTPRequestMethod}
	}
}

// blocking function
func (s *Server) Run() {
	router := mux.NewRouter()

	/*
		takes path string and
		handle function
		type HandlerFunc func(ResponseWriter, *Request)
	*/
	/*ServeHTTP(ResponseWriter, *Request)*/
	router.HandleFunc(rout0, makeHTTPHandleFunc(HandleAccount))
	/* ListenAndServe(addr string, handler Handler) error */
	log.Printf("Listern Server has started on sockaddr=%s", s.Sockaddr)
	startup_error := http.ListenAndServe(s.Sockaddr, router)
	if startup_error != nil {
		log.Printf("error: %s\nin %s", startup_error, "Server/Run/http.ListenAndServe")
	}

}
