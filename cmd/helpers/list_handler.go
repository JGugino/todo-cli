package helpers

type ListItem struct{
	Id string
	Name string
	Completed bool
	ListId string
}

type TodoList struct{
	Id string
	Name string
	ListItems []ListItem
}

type ListHandler struct{
	Lists []TodoList
	Items []ListItem
}

