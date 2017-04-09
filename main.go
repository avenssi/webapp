package main 

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is the first page on aliyun</h1>")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}