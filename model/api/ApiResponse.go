package api

/*
Describes all API Responses Structures

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






type AddBookCategoryResponse struct {
	categoryId int
	parentId int
	categoryName string
}


type RemoveBooKCategoryResponse struct {
	result bool
}