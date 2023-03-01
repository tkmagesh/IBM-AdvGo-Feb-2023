package main

import (
	"fmt"
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

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/customers", customersHandler)
	http.ListenAndServe(":8080", nil)
}
