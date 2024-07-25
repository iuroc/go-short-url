package rule

import "net/http"

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/create", createHandlerFunc)
	router.HandleFunc("/delete", deleteHandlerFunc)
	router.HandleFunc("/search", searchHandlerFunc)
	return router
}
