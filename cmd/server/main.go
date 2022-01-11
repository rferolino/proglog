package main

import (
	"log"

	"github.com/rferolino/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListernAndServe())
}
