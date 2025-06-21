package main

import (
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Millisecond)
		ch <- "ok"
	}()
	w.Write([]byte(<-ch))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
