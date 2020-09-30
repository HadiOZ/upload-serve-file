package main

import (
	"UploadAndServingFile/handler"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hallo from server"))
	})

	http.HandleFunc("/upload", handler.Upload)
	http.HandleFunc("/multiup", handler.MultiUpload)

	log.Printf("server run at %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
