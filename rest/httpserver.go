package rest

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/ping", PingPostman)

	fmt.Println("HTTP server started...")
	log.Fatal(http.ListenAndServe(port, mux))
}
