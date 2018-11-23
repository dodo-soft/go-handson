package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Todo struct {
		ID   string `json:"id"`
		Done bool   `json:"done"`
		Text string `json:"text"`
	}
	TodoList struct {
		List      []*Todo
		currentID int64
	}
)

func NewTodoList() *TodoList {
	return &TodoList{
		List:      make([]*Todo, 0),
		currentID: 0,
	}
}

func main() {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	todoList := NewTodoList()

	// ルーティング定義
	// e.POST("", ...)

	e.Start(":8080")
}
