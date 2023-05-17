package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello EurOpen! %s</h1>\n", hostname)
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on 0.0.0.0:8000, http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
