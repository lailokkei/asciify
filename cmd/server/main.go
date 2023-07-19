package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8080"
	publicDir := "static"
	fs := http.FileServer(http.Dir(publicDir))
	fmt.Printf("hi\n")
	http.ListenAndServe(":"+port, fs)

}
