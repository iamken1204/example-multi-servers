package news

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("news controller at 8001: index page"))
}

func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("news controller at 8001: foo"))
}

func Bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("news controller at 8001: bar"))
}
