package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goCleanArch/internal/model"
	"goCleanArch/internal/repository"
	"net/http"
)

type server struct {
	router *mux.Router
	store  repository.Repository
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.ServeHTTP(writer, request)
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
	s.router.HandleFunc("/user", s.handlCreateuser()).Methods("POST")

}

func (s *server) handlCreateuser() http.HandlerFunc {
	type tmpu struct {
		ID        int    `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Age       uint   `json:"age"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpr := &tmpu{}
		if err := json.NewDecoder(r.Body).Decode(tmpr); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		tmp := model.NewUser()
		tmp = (*model.User)(tmpr)
		if err := s.store.Create(tmp); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusOK, tmp)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
