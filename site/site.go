package site

import "net/http"

func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("site controller at 8002: foo"))
}

func Bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("site controller at 8002: bar"))
}
