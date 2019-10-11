package errorhandler

import "net/http"

// BadRequest writes a bad request to the response
func BadRequest(w http.ResponseWriter) {
	http.Error(w, "Bad method", http.StatusBadRequest)
}
