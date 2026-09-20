package main

import (
	"errors"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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

func (todos *Todos) toggle(index int) error {
	//check index is valid first!
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index] //grab todo
	todo.Completed = !todo.Completed //flip complete

	if todo.Completed {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title //you could save (*todos) to a var t := *todos
	return nil
}


func (todos *Todos) print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "X"
		completedAt := ""

		if t.Completed {
			completed = "DONE"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		tbl.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	tbl.Render()
}

