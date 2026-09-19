package main

impoert (
	"time"
)

type Todo struct{
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time //* lets this be nil 
}
