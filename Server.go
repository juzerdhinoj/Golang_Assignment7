package main

import (
	"fmt"
	"net/http"
)

func DirServer() {

	dir := http.Dir("../Server")
	server := http.FileServer(dir)
	http.Handle("/", server)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Serving")

}
