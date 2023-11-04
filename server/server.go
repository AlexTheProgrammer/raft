package server

import (
	"embed"
	"log"
	"net/http"

	"github.com/AlexTheProgrammer/raft/build"
)

var a embed.FS

func Serve(assets embed.FS) {
	a = assets
	port := ":3000"
	log.Printf("listening on port %v", port)
	handler := middleware(http.FileServer(http.Dir("./dist")))
	log.Fatal(http.ListenAndServe(port, handler))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform actions before handling the request, e.g., authentication, logging, etc.

		build.Build(a)

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)

		// Perform actions after handling the request, if necessary.
	})
}
