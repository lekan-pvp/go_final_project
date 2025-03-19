package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"go1f/pkg/api"
)

func Run(webDir string) error {
	api.Init(`20060102`)
	port := 7540
	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		if eport, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(eport)
		}
	}
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
