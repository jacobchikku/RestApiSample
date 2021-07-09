package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func CreateDemoDetail(response http.ResponseWriter, request *http.Request) {
	var tpl *template.Template
	//response.Header().Set("content-type", "application/json")
	var dm DemoDetails
	//_ = json.NewDecoder(request.Body).Decode(&dm)
	dm.Name = request.FormValue("txtname")
	dm.Email = request.FormValue("txtemail")
	dm.Phone = request.FormValue("txtphone")
	dm.Product = request.FormValue("txtproduct")
	dm.Message = request.FormValue("txtmessage")

	collection := client.Database("DBDemo").Collection("DemoDetails")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, dm)
	fmt.Println(result)

	//json.NewEncoder(response).Encode(result)
	//Show home pages with Data updated status
	results := Resp{Respdata: "Booking Added Successfully. One of our team members will be in touch with you soon."}
	outres := []Resp{results}

	response.Header().Set("Content-Type", "text/html")
	tpl, _ = tpl.ParseFiles("static/resp.html") // Parse template file.
	tpl.Execute(response, outres)
}
