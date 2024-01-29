package rest

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("%s: [%s] - %s ", time.Now().Format(time.RFC3339), r.Method, r.RequestURI)
		log.WithFields(log.Fields{
			"request": r.Method,
			"uri":     r.RequestURI,
		}).Info()
		next.ServeHTTP(w, r)
	})
}
