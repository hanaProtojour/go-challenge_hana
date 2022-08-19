package server

import (
	"go-challenge_hana/internal/hash"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
}

func (s *httpServer) HandlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	hash.DecodeValidateHashJSON(w, r)

}

func NewhttpServer(addr string) *http.Server {
	server := &httpServer{}
	r := mux.NewRouter()
	r.HandleFunc("/", server.HandlePost).Methods("POST")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
