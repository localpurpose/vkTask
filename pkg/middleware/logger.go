package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Println(fmt.Sprintf("%q", x))
		fn(w, r)
	}
}
