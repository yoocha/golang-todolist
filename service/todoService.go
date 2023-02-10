package service

import (
	entity "example.com/todolist/entity"
)

type TodoService struct {
	repository *entity.TodoRepository
}

func CreateTodoService() *TodoService {
	return &TodoService{repository: entity.CreateTodoRepository()}
}

func (s *TodoService) GetTodos() []entity.Todo {
	return s.repository.FindAll()
}

func (s *TodoService) GetTodo(id int64) entity.Todo {
	return s.repository.FindById(id)
}

func (s *TodoService) SaveTodo(todo entity.Todo) entity.Todo {
	return s.repository.Save(todo)
}

func (s *TodoService) UpdateTodo(id int64, todo entity.Todo) entity.Todo {
	entity := s.repository.FindById(id)

	entity.UpdateTodo(todo.Todo)

	return s.repository.Update(entity)
}

func (s *TodoService) CompleteTodo(id int64) entity.Todo {
	entity := s.repository.FindById(id)

	entity.Complete()

	return s.repository.Update(entity)
}

func (s *TodoService) DeleteTodo(id int64) {
	s.repository.DeleteById(id)
}
