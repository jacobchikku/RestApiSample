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

func UpdateDemoDetails(response http.ResponseWriter, request *http.Request) {

	var tpl *template.Template
	fmt.Println("....... Init update function....")
	var filter interface{}
	objid := request.FormValue("txtid")
	objid = strings.Replace(objid, "\"", "\"", 2)

	// replace brackets ()  and double quotes
	str := string("ObjectID(")
	str2 := string(")")
	out := strings.TrimLeft(strings.TrimRight(objid, str2), str)
	//fmt.Println(out)
	objid = out
	objid = strings.Trim(objid, "\"")
	fmt.Println(objid)
	// get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer close(client, ctx, cancel)

	// filter object is used to select a single
	// document matching that matches.
	docID, err := primitive.ObjectIDFromHex(objid)
	if err != nil {
		log.Fatal(err)
	}

	filter = bson.D{{"_id", bson.D{{"$eq", docID}}}}

	// The field of the document that need to updated.
	update := bson.D{
		{"$set", bson.D{
			{"name", request.FormValue("txtname")},
			{"email", request.FormValue("txtemail")},
			{"phone", request.FormValue("txtphone")},
			{"product", request.FormValue("txtproduct")},
			{"message", request.FormValue("txtmessage")},
		}},
	}

	// Returns result of updated document and a error.
	result, err := UpdateOne(client, ctx, "DBDemo",
		"DemoDetails", filter, update)

	// handle error
	if err != nil {
		panic(err)
	}

	// print count of documents that affected
	fmt.Println("update single document")
	fmt.Println(result.ModifiedCount)

	//Show home pages with Data updated status
	results := Resp{Respdata: "Data updated Successfully."}
	outres := []Resp{results}

	response.Header().Set("Content-Type", "text/html")
	tpl, _ = tpl.ParseFiles("static/resp.html") // Parse template file.
	tpl.Execute(response, outres)
}
