package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/root/projectGo/static"))
	fmt.Println(fs)
	http.Handle("/", fs)

	log.Println("Listening on :8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
