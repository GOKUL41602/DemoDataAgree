package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection = db().Database("goTest").Collection("users")

func main() {
	createProfile()
	deleteProfile("ananth")
	getAllUsers()
}

type user struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}

func createProfile() {

	// for adding Content-type

	var person user
	person.Name = "Gokul"
	person.Age = 21
	person.City = "Harur"
	// storing in person variable of type user

	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)

}

func getAllUsers() {

	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	fmt.Println(results)
}

func deleteProfile(name string) {

	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase

	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"name", name}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)

}
