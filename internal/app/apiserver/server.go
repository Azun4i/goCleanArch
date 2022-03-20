package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"goCleanArch/internal/model"
	"goCleanArch/internal/usecases"
	"net/http"
)

type server struct {
	router *mux.Router
	//store  repository.Repository
	logic usecases.UseCaseLogic
}

type tmpu struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       string `json:"age"`
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

func Newserver(logic usecases.UseCaseLogic) *server {
	s := &server{
		logic:  logic,
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/user", s.handlCreateuser()).Methods("POST")
	s.router.HandleFunc("/user/{id}", s.handlGetUserbyId()).Methods("GET")
	s.router.HandleFunc("/user/{id}", s.handlEditUser()).Methods("PUT")
	s.router.HandleFunc("/user/{id}", s.handlDeleteUser()).Methods("DELETE")
}

//handlCreateuser
// create user
func (s *server) handlCreateuser() http.HandlerFunc {
	tmpr := &tmpu{}
	return func(w http.ResponseWriter, r *http.Request) {

		if err := json.NewDecoder(r.Body).Decode(tmpr); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		tmp := model.NewUser()
		tmp.ID = "1" // tests
		tmp.Lastname = tmpr.Lastname
		tmp.Firstname = tmpr.Firstname
		tmp.Email = tmpr.Email
		tmp.Age = tmpr.Age
		if err := s.logic.Create(tmp); err != nil {
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
			s.error(w, r, http.StatusBadRequest, nil)
		}

		u, err := s.logic.FindById(tmp.ID)
		if err != nil {
			s.error(w, r, http.StatusNoContent, errors.New("user not found"))
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

// handlEditUser edit user by id
func (s *server) handlEditUser() http.HandlerFunc {

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
		u.Age = tmpU.Age

		err := s.logic.Edit(u)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

//handlDeleteUser delete user by id
func (s *server) handlDeleteUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		p := mux.Vars(r)
		id, ok := p["id"]
		fmt.Println(id)
		if !ok {
			s.error(w, r, http.StatusBadRequest, errors.New("PPC"))
		}

		if err := s.logic.Delete(id); err != nil {
			s.error(w, r, http.StatusNoContent, err)
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
