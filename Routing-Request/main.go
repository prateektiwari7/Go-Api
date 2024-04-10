package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", hellowhandler).Methods(http.MethodGet)
	r.HandleFunc("/greet/{name}", greethandler).Methods(http.MethodGet)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func hellowhandler(r http.ResponseWriter, w *http.Request) {
	fmt.Fprintln(r, "helloHandler GET Method Called")
}

func greethandler(r http.ResponseWriter, w *http.Request) {
	fmt.Fprintln(r, "Greet Get Method Call")
}
