package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Que dice la banda %s", "Bien o kellogs?")

	})

	srv := http.Server{
		Addr: ":8080",
	}
	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
