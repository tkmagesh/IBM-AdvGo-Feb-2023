package main

import (
	"fmt"
	"log"
	"net/http"
)

/* handlers */
func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

func productsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(res, "All the product info will be served")
	case "POST":
		fmt.Fprintf(res, "The given new product will be added")
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func customersHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "All customers requests will be processed")
}

/* Server implementation */
type AppServer struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func (appServer *AppServer) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.handlers[pattern] = handler
}

/* http.Handler interface implementation */
func (appServer *AppServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s - %s\n", req.Method, req.URL.Path)
	resourceRequested := req.URL.Path
	if handler, exists := appServer.handlers[resourceRequested]; exists {
		handler(res, req)
		return
	}
	http.Error(res, "Resource not found", http.StatusNotFound)
}

func NewAppServer() *AppServer {
	return &AppServer{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func main() {
	appServer := NewAppServer()
	appServer.Register("/", indexHandler)
	appServer.Register("/products", productsHandler)
	appServer.Register("/customers", customersHandler)
	http.ListenAndServe(":8080", appServer)
}
