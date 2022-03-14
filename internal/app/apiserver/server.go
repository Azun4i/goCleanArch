package apiserver

import (
	"github.com/gorilla/mux"
	"goCleanArch/internal/repository"
)

type server struct {
	mux   *mux.Router
	store repository.Store
}

func newserver(store repository.Store) *server {
	s := &server{
		store: store,
		mux:   mux.NewRouter(),
	}

	return s
}
