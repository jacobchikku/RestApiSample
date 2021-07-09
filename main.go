package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomePageHandler).Methods("GET")
	router.HandleFunc("/adddemo", CreateDemoDetail).Methods("POST")
	router.HandleFunc("/getdemo", GetAllDemoDetails).Methods("GET")
	router.HandleFunc("/viewdemo/{id}", GetDemoDetail).Methods("GET")
	router.HandleFunc("/updatedemo", UpdateDemoDetails).Methods("POST")
	router.HandleFunc("/deletedemo", DeleteDemo).Methods("POST")
	//router.HandleFunc("/updatedemo/{id}", updatePost).Methods("PUT")

	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":8080", router))

}

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	var tpl *template.Template
	response.Header().Set("Content-Type", "text/html")
	tpl, _ = tpl.ParseFiles("static/index.html") // Parse template file.
	tpl.Execute(response, "")

}
