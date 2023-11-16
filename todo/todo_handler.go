package todo

import (
	"encoding/json"
	"errors"
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

func MustLoadTodoList() (TodoList, error){
	fileData, err := ReadWholeFile("./", "data.json")

	if err != nil{
		fileContents := TodoList{
			ListName: "Default List",
			TodoItems: make(map[string]TodoItem),
		}

		fileContents.TodoItems["Example One"] = TodoItem{TodoName: "Example One", TodoValue: "This is the first example todo item.", Completed: false}
		fileContents.TodoItems["Example Two"] = TodoItem{TodoName: "Example Two", TodoValue: "This is the second example todo item.", Completed: true}

		marshalledData, err := json.Marshal(fileContents)

		if err != nil{
			panic("Unable to marshal default list contents")
		}

		CreateNewFile("./", "data.json", string(marshalledData))
	}

	var list TodoList

	mErr := json.Unmarshal(fileData, &list)

	if mErr != nil {
		panic("Unable to marshal")
	}

	return list, nil
}

func (handler *TodoHandler) AddList(name string, items map[string]TodoItem){
	handler.TodoLists[name] = &TodoList{ListName: name, TodoItems: items};
}

func (handler *TodoHandler) AddTodo(list string, name string, value string, completed bool) error{
	exists := handler.TodoExists(list, name)

	if exists {
		return errors.New("Todo named (" + name + ") already exists")
	}

	handler.TodoLists[list].TodoItems[name] = TodoItem{TodoName: name, TodoValue: value, Completed: completed}

	return nil
}

func (handler *TodoHandler) RemoveTodo(list  string, name string) error{
	exists := handler.TodoExists(list, name)

	if !exists {
		return errors.New("No todo named (" + name + ") found")
	}

	delete(handler.TodoLists[list].TodoItems, name)

	return nil
}

func (handler *TodoHandler) CompleteTodo(list string, name string) error{
	exists := handler.TodoExists(list, name)

	if !exists {
		return errors.New("No todo named (" + name + ") found")
	}

	handler.TodoLists[list].TodoItems[name] = TodoItem{
		TodoName: handler.TodoLists[list].TodoItems[name].TodoName, 
		TodoValue: handler.TodoLists[list].TodoItems[name].TodoValue, 
		Completed: true,
	}

	return nil
}


func (handler *TodoHandler) TodoExists(list string, name string) bool{
	if _, ok := handler.TodoLists[list].TodoItems[name]; !ok{
		return false
	}

	return true
}

func (handler *TodoHandler) ListTodoItemsFromSpecifiedList(list string) map[string]TodoItem{
	return handler.TodoLists[list].TodoItems
}