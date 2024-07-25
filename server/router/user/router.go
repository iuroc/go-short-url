package user

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/login", loginHandlerFunc)
	router.HandleFunc("/register", registerHandlerFunc)
	router.HandleFunc("/logout", logoutHandlerFunc)
	return router
}
