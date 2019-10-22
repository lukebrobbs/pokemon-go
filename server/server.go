package server

import (
	"fmt"
	"net/http"

	"github.com/lukebrobbs/pokemonServer/pokemon"
)

// Start creates Pokemon routes and starts up the server at the given port.
func Start(port string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
	})
	http.Handle("/pokemon/{pokemon}", http.HandlerFunc(pokemon.Finder))

	http.ListenAndServe(":"+port, nil)
}
