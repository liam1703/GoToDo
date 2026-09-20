package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo item") //The use of & is like saying use the actual thing not a copy
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo item")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo item (format: index:title)")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo item")
	flag.BoolVar(&cf.List, "list", false, "List all todo items")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) error {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		return todos.delete(cf.Del)
	case cf.Edit != "":
		index, title, err := parseEdit(cf.Edit)
		if err != nil {
			return err
		}
		return todos.edit(index, title)
	case cf.Toggle != -1:
		return todos.toggle(cf.Toggle)
	default:
		return errors.New("no valid command provided, use -h for help")
	}
	return nil
}

func parseEdit(s string) (int, string, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return 0, "", errors.New("invalid edit format, use index:title")
	}
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("invalid index %q", parts[0])
	}
	return index, parts[1], nil
}
