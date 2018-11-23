package main

import (
	"errors"
	"fmt"

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

func (t *TodoList) Add(text string) *Todo {
	todo := &Todo{
		ID:   fmt.Sprint(t.currentID),
		Done: false,
		Text: text,
	}
	t.currentID++
	t.List = append(t.List, todo)
	return todo
}

func (t *TodoList) Remove(id string) *Todo {
	for i, todo := range t.List {
		if todo.ID == id {
			t.List = append(t.List[0:i], t.List[i+1:]...)
			return todo
		}
	}
	return nil
}

func (t *TodoList) Find(id string) *Todo {
	for _, todo := range t.List {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

func (t *TodoList) UpdateText(id string, text string) error {
	todo := t.Find(id)
	if todo == nil {
		return errors.New("No such TODO")
	}
	todo.Text = text
	return nil
}

func (t *TodoList) Done(id string, done bool) error {
	todo := t.Find(id)
	if todo == nil {
		return errors.New("No such TODO")
	}
	todo.Done = done
	return nil
}

func main() {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	todoList := NewTodoList()

	e.POST("/todolist/", func(c echo.Context) error {
		text := c.FormValue("text")
		todo := todoList.Add(text)
		return c.JSON(200, todo)
	})

	e.GET("/todolist/", func(c echo.Context) error {
		return c.JSON(200, todoList.List)
	})

	e.GET("/todolist/:id/", func(c echo.Context) error {
		id := c.Param("id")
		todo := todoList.Find(id)
		if todo == nil {
			return c.String(404, "No such TODO")
		}
		return c.JSON(200, todo)
	})

	e.POST("/todolist/:id/done/", func(c echo.Context) error {
		id := c.Param("id")
		todo := todoList.Find(id)
		if todo == nil {
			return c.String(404, "No such TODO")
		}
		todo.Done = true
		return c.JSON(200, todo)
	})

	e.POST("/todolist/:id/undone/", func(c echo.Context) error {
		id := c.Param("id")
		todo := todoList.Find(id)
		if todo == nil {
			return c.String(404, "No such TODO")
		}
		todo.Done = false
		return c.JSON(200, todo)
	})

	e.POST("/todolist/:id/text/", func(c echo.Context) error {
		id := c.Param("id")
		text := c.FormValue("text")
		todo := todoList.Find(id)
		if todo == nil {
			return c.String(404, "No such TODO")
		}
		todo.Text = text
		return c.JSON(200, todo)
	})

	e.DELETE("/todolist/:id/", func(c echo.Context) error {
		id := c.Param("id")
		removed := todoList.Remove(id)
		if removed == nil {
			return c.String(404, "No such TODO")
		}
		return c.JSON(200, removed)
	})

	e.Start(":8080")
}
