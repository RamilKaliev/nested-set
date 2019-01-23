package service

/*
RestService - RestAPI registers endpoint
 */

import (
	"books-store-api/handler"
	"log"
	"net/http"
)

func Init()  {

	log.Println("RestService is Starting")

	// register endpoints and their handlers
	// add node
	http.HandleFunc("/books/categories/add", handler.AddBookCategory)
	// remove node
	http.HandleFunc("/books/categories/remove", handler.RemoveBookCategory)
	// move node
	http.HandleFunc("/books/categories/move", handler.MoveBookCategory)

	// get node parents
	http.HandleFunc("/books/categories/parents", handler.GetNodeParents)
	// get node children
	http.HandleFunc("/books/categories/children", handler.GetNodeChildren)

	// TODO:
	/*
	- Получить всех соседей узла с общим родителем (id[])
	- Получить уровень вложенности узла (int)
	 */

	// open socket
	http.ListenAndServe("localhost:8080", nil)
}