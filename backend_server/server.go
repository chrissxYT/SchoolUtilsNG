package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bfr := []byte(r.URL.Path)
			fn := fmt.Sprintf("%8x", rand.Int31())
			for _, err := os.Stat(fn); !os.IsNotExist(err); {
				fn = fmt.Sprintf("%8x", rand.Int31())
			}
			err := ioutil.WriteFile(fn, bfr, 0644)
			if err != nil {
				log.Print(err)
				fmt.Fprintf(w, "An error occured: %v", err)
				return
			}
			log.Printf("Saved file %s successfully.", fn)
			fmt.Fprintf(w, "Saved file %s successfully.", fn)
		} else {
			fmt.Fprintf(w, "This is an empty page.")
		}
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
