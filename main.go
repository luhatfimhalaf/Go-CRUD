package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server running on port 8080") //untuk memberikan log websitenya
	http.ListenAndServe(":8080",nil) //kalau untuk golang kita perlu membawa package bawaan yaitu http kemudian kita "ListenAndServe(address string,Handler http)"
}