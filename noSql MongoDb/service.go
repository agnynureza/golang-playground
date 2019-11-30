package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func connect() (*mongo.Database, error) {
	fmt.Println("connection to mongo")
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://tunaiku:tunaiku2019@ds149218.mlab.com:49218/tunaiku-testing").SetRetryWrites(false)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("tunaiku-testing"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "wick"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student, 0)
	for csr.Next(ctx) {
		var row student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	if len(result) > 0 {
		for i, k := range result {
			fmt.Println(k)
			fmt.Println("Name  :", result[i].Name)
			fmt.Println("Grade :", result[i].Grade)
		}
	}
}

func update() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	selector := bson.M{"name": "wick"}
	changes := student{"John Wick", 10}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func remove() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").DeleteOne(ctx, bson.M{"name": "wick"})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}

//aggregate data

func aggregateData() {
	pipeline := make([]bson.M, 0)
	err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`[
			{
				"$group":{
					"_id": null,
					"Total" :{"$sum": 1}
				}
			},{
				"$project":{
					"Total": 1,
					"_id": 0
				}
			}
		]`)), true, &pipeline)

	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)
	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}
}
