package server

import (
	"go_challenge/internal/hash"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
}

func (s *httpServer) handlePost(w http.ResponseWriter, r *http.Request) {
	hash.DecodeValidateHashJSON(w, r)

}

func NewhttpServer(addr string) *http.Server {
	server := &httpServer{}
	r := mux.NewRouter()
	r.HandleFunc("/", server.handlePost).Methods("POST")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
