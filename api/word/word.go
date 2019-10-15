package word

import (
	"fmt"
	"net/http"
)

type WordAPI struct{}

func (word WordAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		DoPost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Unsupported method '%v' to %v \n", r.Method, r.URL)
	}
}