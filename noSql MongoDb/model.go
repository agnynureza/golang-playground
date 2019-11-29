package main

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"grade"`
}
