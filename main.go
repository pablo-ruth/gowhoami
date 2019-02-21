package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, "Failed to get hostname: %s", err)
		}

		fmt.Fprintf(w, "Welcome on %s !!!", hostname)
	})

	http.ListenAndServe(":3333", nil)
}
