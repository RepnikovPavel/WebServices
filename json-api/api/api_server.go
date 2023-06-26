package api

import (
	"encoding/json"
	"fmt"
	"json-api/DataStructures"
	"json-api/Storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	urd UserRegistrationDesk
	ugs UserGetterSetter
)

type UserRegistrationDesk struct {
	// link to db
	// mutex
}

// TODO: add custom error with db error code
func (rd *UserRegistrationDesk) IsUserExists() (bool, error) {
	return false, nil
}

func (rd *UserRegistrationDesk) RegisterAUser(ucs DataStructures.UserClientSide) error {
	// if user exists
	//    create UserAccount
	//    create UserServerSide
	//    create UserServerPayload
	//    make transactions in db
	//    update cach
	//    using mutex for goroutines
	return nil
}

func (rd *UserRegistrationDesk) DeleteAUser(ucs DataStructures.UserServerSide) error {
	// if user exists
	//    make transactions in db
	//    update cach
	//    using mutex for goroutines
	return nil
}

type UserGetterSetter struct {
}

func (ugs *UserGetterSetter) UpdateAUser(ucs DataStructures.UserClientSide) error {
	// if user exists
	//    update UserServerSide
	//    update UserserverPayload
	//    make transactions in db
	//    update cach
	//    using mutex for goroutines
	return nil
}
func (ugs *UserGetterSetter) ReadAUser() (DataStructures.UserClientSide, error) {
	// if user exists
	//    read UserClientSide
	//    safety read from db
	//    update cach
	//    using mutex for goroutines
	return DataStructures.UserClientSide{}, nil
}

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
	// Header signature is
	// type Header map[string][]string
	// Set(key: x string,value: y string)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
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
	Sockaddr   string
	AccountAPI Storage.AccountAPI
}

func NewServer(conf ServerConfig, AccountAPI Storage.AccountAPI) *Server {
	return &Server{
		Sockaddr:   conf.Sockaddr,
		AccountAPI: AccountAPI,
	}
}

const (
	rout0 string = "/rout0"
	rout1 string = "/rout1"
	rout2 string = "/rout2"
	rout3 string = "/rout3"
	rout4 string = "/rout4"
)

// CRUD for accounts
// functions signature is payload function:
// type payload func(w http.ResponseWriter, r *http.Request) error

func PostAccount(w http.ResponseWriter, r *http.Request) error {
	// blocking op
	// 1) parse body
	// 2) is user exists
	// 3) do payload
	// isexists, err := urd.IsUserExists()

	// 1)
	fmt.Println("POST method works")

	return nil
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("DELETE method works")

	return nil
}

func PutAccount(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("PUT method works")

	return nil
}

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("%s", "GET method works")
	ExStruct := DataStructures.NewExampleStruct()
	return WriteJSON(w, http.StatusOK, ExStruct)
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
	log.Printf("Listern Server has started on sockaddr %s", s.Sockaddr)
	fmt.Printf("supported routes:\n%s\n",
		rout0)
	startup_error := http.ListenAndServe(s.Sockaddr, router)
	if startup_error != nil {
		log.Printf("error: %s\nin %s", startup_error, "Server/Run/http.ListenAndServe")
	}
}
