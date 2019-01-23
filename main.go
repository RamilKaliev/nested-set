package main

import (
	"books-store-api/repository"
	"books-store-api/service"
)

/*
Application starting point

- initialization of all required services and handlers.
 */
func main() {

	// init postgres repository
	repository.Init()

	// init rest service: handle endpoints with functions
	service.Init()

}