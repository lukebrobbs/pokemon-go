package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukebrobbs/pokemon-go/pokemon"
)

// Start creates Pokemon routes and starts up the server at the given port.
func Start(port string) {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
	})
	r.Handle("/pokemon/{pokemon}", http.HandlerFunc(pokemon.Finder))
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
