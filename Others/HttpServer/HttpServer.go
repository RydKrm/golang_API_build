package main

import (
	"fmt"
	"net/http"
)

func homePage(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to http server");
	fs := http.FileServer(http.Dir("static/"));
	http.Handle("/static/", http.StripPrefix("/static/",fs))
	http.ListenAndServe(":8080",nil)
	})
}

func main() {
	fmt.Println("Process dynamic requests");
	homePage()
}