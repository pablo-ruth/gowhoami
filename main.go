package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = 1

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, "Failed to get hostname: %s", err)
		}

		fmt.Fprintf(w, "Welcome on %s !!! version %d)", hostname, version)
	})

	fmt.Println("Listening on :3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
