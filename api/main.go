package main

import (
	"net/http"
)

func myapi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Merhaba Yabanci"))
}

func main() {
	http.HandleFunc("/myapi", myapi)
	http.ListenAndServe(":3030", nil)
}
