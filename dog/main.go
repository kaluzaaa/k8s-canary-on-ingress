package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tomasen/realip"
)

func pet(w http.ResponseWriter, r *http.Request) {
	clientIP := realip.FromRequest(r)
	log.Println("GET /pet from", clientIP)
	message := "Dogs are always happy to see you!"
	w.Write([]byte(message))
}
func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Server is starting...")
	http.HandleFunc("/", pet)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
