package server

import (
	"net/http"

	"github.com/ryzaer/kyrafs/internal/handler"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/put", handler.Put)

	return &Server{
		mux: mux,
	}
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
