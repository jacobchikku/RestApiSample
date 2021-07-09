package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type DemoDetails struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Product string             `json:"product,omitempty" bson:"product,omitempty"`
	Message string             `json:"message,omitempty" bson:"message,omitempty"`
}
type Resp struct {
	Respdata string
}

func init() {
	fmt.Println("Init the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
}

// UpdateOne is a user defined method, that update
// a single document matching the filter.
// This methos accepts client, context, databse,
// collection, filter and update filter and update
// is of type interface this method returns
// UpdateResult and an error if any.
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the databse and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the
	// filter will get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}

// query method returns a cursor and error.
func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
