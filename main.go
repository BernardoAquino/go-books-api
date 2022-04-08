package main

import (
	"github.com/bernardoaquino/go-books-api/database"
	"github.com/bernardoaquino/go-books-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
