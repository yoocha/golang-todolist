package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Todo string
	Done bool `gorm:"default:false"`
}

func (e *Todo) Complete() {
	e.Done = true
}

func (e *Todo) UpdateTodo(todo string) {
	e.Todo = todo
}
