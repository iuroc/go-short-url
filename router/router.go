package router

import "net/http"

func MainRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})
	router.Handle("/api/user/", http.StripPrefix("/api/user", UserRouter()))
	router.Handle("/api/url/", http.StripPrefix("/api/url", URLRouter()))
	return router
}
