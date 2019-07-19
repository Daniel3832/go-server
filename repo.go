package main

import "fmt"

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// find todo and return this todo object, return empty Todo if not found
func RepoFindTodo(id int) Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

// create a new todo, id increment
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

// remove a todo from todos list. else if not found this id, fmt error
func RepoDestroyTodo(id int) error {
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
