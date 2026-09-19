package main

import (
	"errors"
	"slices"
	"time"
)

type Todo struct{
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time //* lets this be nil 
}

type Todos []Todo

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("invalid index")
	}
	return nil
}

func (todos *Todos) Add(title string) { //(todos *Todos) is a pointer receiver, so we can modify the original slice
	todo := Todo{
		Title:     title,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = slices.Delete(*todos, index, index+1) // Delete can be used on slices 
	return nil
}