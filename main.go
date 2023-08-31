package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, "Hello World")

		if err != nil {
			fmt.Println("Error writing to response.")
		}

		fmt.Printf("%d", n)
	})

	http.ListenAndServe(":3000", nil)
}
