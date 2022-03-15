package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"goCleanArch/internal/model"
	"goCleanArch/internal/repository"
	"net/http"
	"strconv"
)

type server struct {
	router *mux.Router
	store  repository.Repository
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func newserver(store repository.Repository) *server {
	s := &server{
		store:  store,
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/user", s.handlCreateuser()).Methods("POST")      // work
	s.router.HandleFunc("/user/{id}", s.handlGetUserbyId()).Methods("GET") // work
	s.router.HandleFunc("/user/{id}", s.handlEditUser()).Methods("PUT")
}

//handlCreateuser
// create user
func (s *server) handlCreateuser() http.HandlerFunc {

	type tmpu struct {
		ID        string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Age       string `json:"age"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpr := &tmpu{}
		if err := json.NewDecoder(r.Body).Decode(tmpr); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		fmt.Printf("%#v\n", tmpr)
		tmp := model.NewUser()
		tmp.ID = uuid.New().String()
		tmp.Lastname = tmpr.Lastname
		tmp.Firstname = tmpr.Firstname
		tmp.Email = tmpr.Email
		tmp.Age, _ = strconv.Atoi(tmpr.Age)

		fmt.Printf("%#v\n", tmp)
		if err := s.store.Create(tmp); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusOK, tmp)
	}
}

//handlGetUserbyId get user
func (s *server) handlGetUserbyId() http.HandlerFunc {
	type tmp_req struct {
		ID string
	}
	var ok bool

	tmp := &tmp_req{}
	return func(w http.ResponseWriter, r *http.Request) {
		p := mux.Vars(r)

		tmp.ID, ok = p["id"]
		fmt.Println(tmp.ID)
		if !ok {
			s.error(w, r, http.StatusNoContent, nil)
		}

		u, err := s.store.FindById(tmp.ID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, nil)
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handlEditUser() http.HandlerFunc {
	type tmpu struct {
		ID        string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Age       string `json:"age"`
	}

	tmpU := &tmpu{}
	return func(w http.ResponseWriter, r *http.Request) {
		p := mux.Vars(r)
		id, ok := p["id"]
		if !ok {
			s.error(w, r, http.StatusBadRequest, nil)

		}

		if err := json.NewDecoder(r.Body).Decode(&tmpU); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		tmpU.ID = id
		u := model.NewUser()

		u.ID = tmpU.ID
		u.Firstname = tmpU.Firstname
		u.Lastname = tmpU.Lastname
		u.Email = tmpU.Email
		u.Age, _ = strconv.Atoi(tmpU.Age)

		err := s.store.Edit(u)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

// error send error to Request
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
