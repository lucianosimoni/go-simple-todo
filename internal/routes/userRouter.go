// User Router
package routes

import (
	"net/http"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am the user."))
}
