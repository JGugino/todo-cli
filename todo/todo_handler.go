package todo

import (
	"encoding/json"
	"fmt"
)

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

func LoadTodoList(){
	fileData, err := ReadWholeFile("./", "data.json")

	if err != nil{
		fileContents := TodoList{
			ListName: "Default List",
			TodoItems: make(map[string]TodoItem),
		}

		fileContents.TodoItems["Example One"] = TodoItem{TodoName: "Example One", TodoValue: "This is the first example todo item.", Completed: false}
		fileContents.TodoItems["Example Two"] = TodoItem{TodoName: "Example Two", TodoValue: "This is the second example todo item.", Completed: true}

		marshalledData, _ := json.Marshal(fileContents)
		CreateNewFile("./", "data.json", string(marshalledData))
		fmt.Println("error: " + err.Error())
	}

	var list TodoList

	mErr := json.Unmarshal(fileData, &list)

	if mErr != nil {
		panic("Unable to marshal")
	}

	fmt.Println(list)
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