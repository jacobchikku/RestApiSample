package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteDemo(response http.ResponseWriter, request *http.Request) {
	var tpl *template.Template
	fmt.Println("....... Init Delete function....")

	objid := request.FormValue("txtid")

	fmt.Println(objid)
	objid = strings.Replace(objid, "\"", "\"", 2)

	// replace brackets ()  and double quotes
	str := string("ObjectID(")
	str2 := string(")")
	out := strings.TrimLeft(strings.TrimRight(objid, str2), str)
	//fmt.Println(out)
	objid = out
	objid = strings.Trim(objid, "\"")
	fmt.Println(objid)
	docID, err := primitive.ObjectIDFromHex(objid)
	if err != nil {
		log.Fatal(err)
	}
	// get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	// free resource when main function is returned
	defer close(client, ctx, cancel)

	// This query delete document when the maths
	// field is greater than 60

	if err != nil {
		log.Fatal(err)
	}

	query := bson.D{{"_id", bson.D{{"$eq", docID}}}}

	// Returns result of deletion and error
	result, err := deleteOne(client, ctx, "DBDemo", "DemoDetails", query)
	if err != nil {
		log.Fatal(err)
	}
	// print the count of affected documents
	fmt.Println("No.of rows affected by DeleteOne()")
	fmt.Println(result.DeletedCount)

	//Show home pages with Data updated status
	results := Resp{Respdata: "Data Deleted Successfully."}
	outres := []Resp{results}

	response.Header().Set("Content-Type", "text/html")
	tpl, _ = tpl.ParseFiles("static/resp.html") // Parse template file.
	tpl.Execute(response, outres)

}
