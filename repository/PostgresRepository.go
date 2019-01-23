package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)


var _connection *sql.DB


func Init() {
	// init database connection pool
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "0.0.0.0", "5432", "postgres", "postgres", "postgres")
	connection, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	connection.SetMaxOpenConns(10)

	_connection = connection
	log.Println("Connected to Postgres Database")
}


/**
Append book category into existed parent node
*/
func AddBookCategory(parentId int, bookCategory string) int {

	rows, err := _connection.Query("SELECT public.f_add_book_category($1, $2);", parentId, bookCategory)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cat_id := -1

	if (rows.Next()) {
		err = rows.Scan(&cat_id)
		if err != nil {
			log.Fatal(err)
		}
	}

	return cat_id

}

/*
Remove book category from
 */
func RemoveBookCategory(categoryId int) bool {

	rows, err := _connection.Query("SELECT public.f_remove_book_category($1);", categoryId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := -1

	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result != -1
}


/*
MoveBookCategory - replace concrete node into another category with all descent nodes.
 */
func MoveBookCategory(categoryId int, oldParentId int, newParentId int) bool {

	rows, err := _connection.Query("SELECT public.f_move_book_category($1, $2, $3);", categoryId, oldParentId, newParentId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := -1

	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result != -1

}



/*
GetNodeParents - return an array of parents' ids for the concrete node.
 */
func GetNodeParents(categoryId int) []int {

	var query = "WITH current_range AS (" +
		"select left_edge, right_edge " +
		"from public.nested_book_categories " +
		"where category_id = $1" +
		") " +
		"select category_id " +
		"from public.nested_book_categories " +
		"where " +
			"left_edge < (select left_edge from current_range limit 1) and " +
			"right_edge > (select right_edge from current_range limit 1) " +
		"order by category_id;"

	rows, err := _connection.Query(query, categoryId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ids []int

	log.Println(rows)

	for (rows.Next()) {
		var id int
		err = rows.Scan(&id)
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	return ids
}


func GetNodeChildren(categoryId int) []int {

	var query = "WITH current_range AS (" +
		"select left_edge, right_edge " +
		"from public.nested_book_categories " +
		"where category_id = $1" +
		") " +
		"select category_id " +
		"from public.nested_book_categories " +
		"where " +
		"left_edge > (select left_edge from current_range limit 1) and " +
		"right_edge < (select right_edge from current_range limit 1) " +
		"order by category_id;"

	rows, err := _connection.Query(query, categoryId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ids []int

	log.Println(rows)

	for (rows.Next()) {
		var id int
		err = rows.Scan(&id)
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	return ids
}