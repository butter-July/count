package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	count := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count += 1
		fmt.Fprintf(w, "%d", count)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
