package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PageHandler(w http.ResponseWriter, r *http.Request){
	// to get the params variable 
	vars := mux.Vars(r);
	pageNumber := vars["title"];

	fmt.Println("Page number ", pageNumber)
}

func main() {
	fmt.Println("Go Routing ");

	route := mux.NewRouter();

	// router handler function 
	route.HandleFunc("/page/{page}", PageHandler);

	http.ListenAndServe(":8000", nil)

}