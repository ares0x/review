package main

import (
	"net/http"
	_ "net/http/pprof"
)


func main() {
	http.HandleFunc("/",Hello)
	http.ListenAndServe(":8080",nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func Map(w http.ResponseWriter, r *http.Request) {
	
}