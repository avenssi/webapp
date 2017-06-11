package main 

import (
    "io"
    "net/http"
    "log"
)


func firstPage(w http.ResponseWriter, r *http.Request) {
	log.Println("New request came in.")
	io.WriteString(w, "<h1>Hello, I am a hero! </h1>")
}


func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil) //:::8000
}
