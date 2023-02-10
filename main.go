package main

import (
	"fmt"
	"github.com/dainoguchi/bstodo-api/config"
	"log"
	"net/http"
)

func main() {
	cfg, _ := config.New()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	log.Printf("listening on port %d", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil); err != nil {
		log.Fatal(err)
	}
}
