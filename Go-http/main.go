package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Hello", helloHandler)
	http.ListenAndServe(":3000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET Method Called")
	case http.MethodPost:
		fmt.Fprintln(w, "Post Method Called")
	case http.MethodPut:
		fmt.Fprintln(w, "Put Method Called")
	default:
		http.Error(w, "Not Catched Method", http.StatusMethodNotAllowed)
	}
}
