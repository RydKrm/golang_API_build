package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1> Getting start with golang  </h1>")
	fmt.Fprint(w, r.URL)
}

func contactHandler(w http.ResponseWriter, r *http.Request) { 
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=mailto:jon@calho </a> </p>") 
}

func main() {
	fmt.Println("Start Go project ");
	http.HandleFunc("/", handleFunc);
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":3000",nil)
}