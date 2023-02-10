package entity

type Repository interface {
	FindAll() []interface{}
	FindById(id int64) interface{}
	Save(entity interface{})
	Update(id int64, entity interface{})
	DeleteById(id int64)
}
