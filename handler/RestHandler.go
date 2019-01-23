package handler

import (
	"books-store-api/model/api"
	"books-store-api/repository"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)


func AddBookCategory(rs http.ResponseWriter, rq *http.Request) {

	var request api.AddBookCategoryRequest

	body, err := ioutil.ReadAll(io.LimitReader(rq.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := rq.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode(err); err != nil {
			panic(err)
		}
	}

	result := repository.AddBookCategory(request.ParentId, request.CategoryName)

	rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rs).Encode(result); err != nil {
		panic(err)
	}

}




func RemoveBookCategory(rs http.ResponseWriter, rq *http.Request) {

	var request api.RemoveBookCategoryRequest

	body, err := ioutil.ReadAll(io.LimitReader(rq.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := rq.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode(err); err != nil {
			panic(err)
		}
	}

	result := repository.RemoveBookCategory(request.CategoryId)

	rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rs).Encode(result); err != nil {
		panic(err)
	}

}



func MoveBookCategory(rs http.ResponseWriter, rq *http.Request) {

	var request api.MoveBookCategoryRequest

	body, err := ioutil.ReadAll(io.LimitReader(rq.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := rq.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode(err); err != nil {
			panic(err)
		}
	}

	result := repository.MoveBookCategory(request.CategoryId, request.OldParentId, request.NewParentId)

	rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rs).Encode(result); err != nil {
		panic(err)
	}
}



func GetNodeParents(rs http.ResponseWriter, rq *http.Request) {

	var categoryIdStr = rq.URL.Query()["categoryId"]

	// if query parameter is missing
	if categoryIdStr == nil {
		panic("Parameter 'categoryId' is missing")
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode("{'message': 'Parameter categoryId is missing'}"); err != nil {
			panic(err)
		}
	}

	// convert categoryId into int
	categoryId, err := strconv.Atoi(categoryIdStr[0])
	if err != nil {
		panic("Parameter value must be integer")
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode("{'message': 'Parameter value must be integer'}"); err != nil {
			panic(err)
		}
	}

	result := repository.GetNodeParents(categoryId)

	rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rs).Encode(result); err != nil {
		panic(err)
	}
}


func GetNodeChildren(rs http.ResponseWriter, rq *http.Request) {

	var categoryIdStr = rq.URL.Query()["categoryId"]

	// if query parameter is missing
	if categoryIdStr == nil {
		panic("Parameter 'categoryId' is missing")
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode("{'message': 'Parameter categoryId is missing'}"); err != nil {
			panic(err)
		}
	}

	// convert categoryId into int
	categoryId, err := strconv.Atoi(categoryIdStr[0])
	if err != nil {
		panic("Parameter value must be integer")
		rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rs.WriteHeader(400) // unprocessable entity
		if err := json.NewEncoder(rs).Encode("{'message': 'Parameter value must be integer'}"); err != nil {
			panic(err)
		}
	}

	result := repository.GetNodeChildren(categoryId)

	rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rs).Encode(result); err != nil {
		panic(err)
	}
}


// TODO: needs to be done
/*
	- Получить всех соседей узла с общим родителем (id[])
	- Получить уровень вложенности узла (int)
 */