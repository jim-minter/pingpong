package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		log.Print("listening")
		http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("serving %s", r.URL)
		}))

	} else {
		t := time.NewTicker(time.Second)

		for {
			cli := &http.Client{Transport: &http.Transport{}}
			go func() {
				resp, err := cli.Get(os.Args[1])
				if err != nil {
					log.Printf("getting %s: %s", os.Args[1], err)
					return
				}
				defer resp.Body.Close()
				log.Printf("getting %s: %s", os.Args[1], resp.Status)
			}()
			<-t.C
		}
	}
}
