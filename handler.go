package greeting

import (
	"fmt"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}