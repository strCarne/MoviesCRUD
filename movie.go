package main

type Movie struct {
	ID string `json: "id"`
	ISBN string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

var movies []Movie