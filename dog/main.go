package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tomasen/realip"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func pet(w http.ResponseWriter, r *http.Request) {
	clientIP := realip.FromRequest(r)
	log.Println("GET /pet from", clientIP)
	message := "Dogs are always happy to see you!"
	enableCors(&w)
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
