package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/geekyharsh05/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// setup database

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("Server started on port http://%s\n", cfg.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed To Start Server")
	}

}
