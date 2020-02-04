package main

import (
	"fmt"
	"log"
	"net/http"
)

func run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bfr := make([]byte, 1024*1024*16)
			n, err := r.Body.Read(bfr)
			if err != nil {
				log.Fatal(err)
			}
			//write bfr to a file
		}
		fmt.Fprintf(w, "This is an empty page.")
	})

	http.ListenAndServe(":4269", nil)

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
