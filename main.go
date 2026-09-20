package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")

	if err := storage.Load(&todos); err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmdFlags := NewCmdFlags()
	if err := cmdFlags.Execute(&todos); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := storage.Save(todos); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}