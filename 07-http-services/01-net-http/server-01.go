package main

import (
	"fmt"
	"log"
	"net/http"
)

type AppServer struct {
}

/* http.Handler interface implementation */
func (appServer *AppServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s - %s\n", req.Method, req.URL.Path)
	switch req.URL.Path {
	case "/":
		fmt.Fprintln(res, "Hello World!")
	case "/products":
		switch req.Method {
		case "GET":
			fmt.Fprintf(res, "All the product info will be served")
		case "POST":
			fmt.Fprintf(res, "The given new product will be added")
		default:
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case "/customers":
		fmt.Fprintln(res, "All customers requests will be processed")
	default:
		http.Error(res, "Resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
