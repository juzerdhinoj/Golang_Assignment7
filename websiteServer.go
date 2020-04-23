package main

import (
	"fmt"
	"net/http"
)

func StaticWebsite() {
	fs := http.FileServer(http.Dir("./website"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":4040", nil)
	if err != nil {
		fmt.Println(err)
	}
}
