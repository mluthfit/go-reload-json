package services

import (
	"fmt"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/favicon.ico", IconHandler)
	http.HandleFunc("/", IndexHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("views"))))

	fmt.Printf("Server started at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
