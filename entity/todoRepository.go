package entity

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func CreateTodoRepository() *TodoRepository {

	user := "root"
	password := "1q2w3e4r5t"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:53306)/todo?charset=utf8&parseTime=true", user, password)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		conn.AutoMigrate(&Todo{})
		repository := TodoRepository{db: conn}
		fmt.Println("DB Connected")
		return &repository
	}
}

func (r *TodoRepository) FindAll() []Todo {
	var todos []Todo
	r.db.Find(&todos)
	return todos
}

func (r *TodoRepository) FindById(id int64) Todo {
	var todo Todo
	r.db.Find(&todo)
	return todo
}

func (r *TodoRepository) Save(entity Todo) Todo {
	r.db.Create(&entity)
	return entity
}

func (r *TodoRepository) Update(entity Todo) Todo {
	r.db.Save(&entity)
	return entity
}

func (r *TodoRepository) DeleteById(id int64) {
	todo := r.db.Find(&Todo{}, id)
	r.db.Delete(&todo)
}
