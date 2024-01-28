package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IfSentient/gomodtest/submod/internal/server"
)

func main() {
	fmt.Println("do thing")
	r := server.CreateServer()

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
