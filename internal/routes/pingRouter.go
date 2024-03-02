// Ping Router
package routes

import (
	"fmt"
	"net/http"
)

func PingRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
