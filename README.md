# GoToDo

A CLI todo list written in Go.

## Usage

```sh
go run . -add "buy milk"
go run . -list
go run . -toggle 0
go run . -edit "0:buy oat milk"
go run . -del 0
```

## Possible extensions

1. Add due dates and priorities so you can sort or filter todos by what matters most.
2. Add tags like "work" or "home" so you can group and filter todos as the list grows.
3. Add an interactive keyboard-driven UI so you can browse and edit todos without typing indices.
