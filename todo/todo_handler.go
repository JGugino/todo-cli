package todo

type TodoHandler struct{
	TodoLists map[string] *TodoList
}

type TodoList struct{
	ListName string `json:"listName"`
	TodoItems map[string]TodoItem `json:"todoItems"`
}

type TodoItem struct{
	TodoName string `json:"todoName"`
	TodoValue string `json:"todoValue"`
	Completed bool `json:"completed"`
}

func NewTodoHandler() *TodoHandler{
	return &TodoHandler{
		TodoLists: make(map[string]*TodoList),
	}
}

func (handler *TodoHandler) AddList(name string, items map[string]TodoItem){
	handler.TodoLists[name] = &TodoList{ListName: name, TodoItems: items};
}

func (handler *TodoHandler) AddTodo(list string, name string, value string, completed bool){
	handler.TodoLists[list].TodoItems[name] = TodoItem{TodoName: name, TodoValue: value, Completed: completed}
}

func (handler *TodoHandler) ListTodoItemsFromSpecifiedList(list string) map[string]TodoItem{
	return handler.TodoLists[list].TodoItems
}