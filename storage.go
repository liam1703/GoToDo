package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

//"Give me a filename and tell me what type of thing this storage will hold. I'll give you back a storage for that type."
// e.g.   NewStorage[int]("nums.json")   = a storage of ints in nums.json
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}


