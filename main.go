package main

import (
	"net/http"

	"github.com/iamken1204/example-multi-servers/news"
	"github.com/iamken1204/example-multi-servers/site"
)

func main() {
	finish := make(chan bool)

	server8001 := http.NewServeMux()
	server8001.HandleFunc("/foo", news.Foo)
	server8001.HandleFunc("/bar", news.Bar)

	server8002 := http.NewServeMux()
	server8002.HandleFunc("/foo", site.Foo)
	server8002.HandleFunc("/bar", site.Bar)

	// The second parameter of http.ListenAndServe
	// is a ServeMux.
	go func() {
		http.ListenAndServe(":8001", server8001)
	}()

	go func() {
		http.ListenAndServe(":8002", server8002)
	}()

	<-finish
}
