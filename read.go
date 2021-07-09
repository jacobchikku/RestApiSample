package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllDemoDetails(response http.ResponseWriter, request *http.Request) {
	fmt.Println("....... Init GetDemoDETAILS function....")
	var tpl *template.Template
	response.Header().Set("content-type", "application/json")
	//var dm1 []DemoDetails
	collection := client.Database("DBDemo").Collection("DemoDetails")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	var dm []DemoDetails
	if err := cursor.All(ctx, &dm); err != nil {

		// handle the error
		panic(err)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	//json.NewEncoder(response).Encode(dm1)
	//tpl.New("some template")                     // Create a template.
	response.Header().Set("Content-Type", "text/html")
	tpl, _ = tpl.ParseFiles("static/view.html") // Parse template file.
	tpl.Execute(response, dm)                   // merge.
	//tpl.ExecuteTemplate(response, "index.html", dm)
}

func GetDemoDetail(response http.ResponseWriter, request *http.Request) {
	//for k, v := range mux.Vars(request) {
	//	fmt.Printf("key=%v, value=%v", k, v)
	//}
	fmt.Println("....... Init GetDemo function....")
	var tpl *template.Template
	params := mux.Vars(request)
	objid := params["id"]
	objid = strings.Replace(objid, "\"", "\"", 2)

	str := string("ObjectID(")
	str2 := string(")")
	out := strings.TrimLeft(strings.TrimRight(objid, str2), str)
	//fmt.Println(out)
	objid = out
	objid = strings.Trim(objid, "\"")
	fmt.Println("Query string key value", objid)

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when mainn dunction is returned
	defer close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter interface{}

	// filter gets all document,
	// with maths field greater that 70
	docID, err := primitive.ObjectIDFromHex(objid)
	if err != nil {
		log.Fatal(err)
	}
	filter = bson.D{{"_id", bson.D{{"$eq", docID}}}}

	// option remove id field from all documents
	//option = bson.D{{"id", 0}}

	// call the query method with client, context,
	// database name, collection name, filter and option
	// This method returns momngo.cursor and error if any.
	cursor, err := query(client, ctx, "DBDemo", "DemoDetails", filter, nil)
	// handle the errors.
	if err != nil {
		panic(err)
	}

	//var results []bson.D

	// to get bson object from cursor,
	// returns error if any.
	var dm []DemoDetails
	if err := cursor.All(ctx, &dm); err != nil {
		// handle the error
		panic(err)
	}
	response.Header().Set("Content-Type", "text/html")
	tpl, _ = tpl.ParseFiles("static/editdemo.html") // Parse template file.
	tpl.Execute(response, dm)                       // merge.

}
