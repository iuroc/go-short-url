package userrouter

import "net/http"

func Router() *http.ServeMux {
	router := http.NewServeMux()
	return router
}
