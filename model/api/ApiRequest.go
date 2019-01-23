package api


/*
Describes all API Requests Structures

- Нормализация дерева (bool)
- Добавление узла в дерево (id)
- Удаление узла из дерева (bool)
- Перенос узла внутри дерева (bool)
- Получить всех родителей узла (id[])
- Получить всех прямых потомков узла (id[])
- Получить всех потомков узла (id[])
- Получить всех соседей узла с общим родителем (id[])
- Получить уровень вложенности узла (int)
*/

/*

 */
type NormalizeBooksCategoriesRequest struct {
	rootId int
}

/*
AddNodeRequest:
- describes set of fields required for adding a new node to the concrete book category.
*/
type AddBookCategoryRequest struct {
	ParentId int
	CategoryName string
}

/*
DeleteBookCategoryRequest:
- describes set of fields required for removing an existing node out of the concrete book category.
 */
type RemoveBookCategoryRequest struct {
	CategoryId int
}

/*
MoveBookCategoryRequest:
- describes set of fields which required for moving on a node to the another book category.
 */
type MoveBookCategoryRequest struct {
	CategoryId int
	OldParentId int
	NewParentId int
}

/*
GetParentsForCategoryRequest:
- set of fields for getting all parent for the current node.
 */
type GetNodeParents struct {
	CategoryId int
}





/*
GetDirectDescendantsForCategoryRequest
- request set for getting direct descendants of the concrete category.
 */
type GetDirectDescendantsForCategoryRequest struct {
	bookCategoryId int
	bookCategoryName string
}

/*
GetSiblingCategoriesForCategoryRequest
- request set for getting sibling categories.
 */
type GetSiblingCategoriesForCategoryRequest struct {
	bookCategoryId int
	bookCategoryName string
}

/*
GetHeightForCategoryRequest
- request set for getting height for the tree part from the concrete category.
*/
type GetHeightForCategoryRequest struct {
	bookCategoryId int
	bookCategoryName string
}