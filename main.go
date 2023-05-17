package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	HOSTNAME, _ := os.Hostname()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		counter, _ := rdb.Incr(ctx, "counter").Result()
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Hello EurOpen! %d %s</h1>\n", counter, HOSTNAME)
		log.Info().
			Str("hostname", HOSTNAME).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Msg(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	log.Info().
		Str("hostname", HOSTNAME).
		Msg("Listening on 0.0.0.0:8000, http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
