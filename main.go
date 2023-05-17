package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	HOSTNAME, _ := os.Hostname()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Hello EurOpen! %s</h1>\n", HOSTNAME)
		log.Info().
			Str("hostname", HOSTNAME).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	})

	log.Info().
		Str("hostname", HOSTNAME).
		Msg("Listening on 0.0.0.0:8000, http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
