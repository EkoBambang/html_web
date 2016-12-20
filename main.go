package main

import "net/http"

func indeks(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}
func about(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "about.html")
}
func base(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "base.html")
}

func main() {
	http.HandleFunc("/", indeks)
	http.HandleFunc("/about", indeks)
	http.HandleFunc("/base", indeks)
	http.ListenAndServe(":8080", nil)
}
