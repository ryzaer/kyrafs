package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("KyraFS Engine Ready"))

	})

	log.Println("KyraFS started at http://127.0.0.1:8000")

	log.Fatal(http.ListenAndServe(":8000", nil))

}