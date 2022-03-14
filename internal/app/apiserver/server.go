package apiserver

import (
	"github.com/gorilla/mux"
	"goCleanArch/internal/repository"
)

type server struct {
	mux   *mux.Router
	store repository.Repository
}

func newserver(store repository.Repository) *server {
	s := &server{
		store: store,
		mux:   mux.NewRouter(),
	}

	///ROUTERI.............

	return s
}
