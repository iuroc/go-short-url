package rule

import "net/http"

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/create", createHandlerFunc)
	return router
}
