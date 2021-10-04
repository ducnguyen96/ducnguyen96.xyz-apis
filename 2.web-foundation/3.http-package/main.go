package main

import (
	"fmt"
	"log"
	"net/http"
)

//type MyMux struct {
//}

//func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path == "/" {
//		sayHelloName(w, r)
//		return
//	}
//	http.NotFound(w, r)
//	return
//}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello myroute!")
	if err != nil {
		return 
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	//mux := &MyMux{}
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}