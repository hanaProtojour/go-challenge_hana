package main

import (
	"go-challenge_hana/internal/server"
)

func main() {
	println("Starting listening on port 3333")
	srv := server.NewhttpServer(":3333")
	srv.ListenAndServe()
}
