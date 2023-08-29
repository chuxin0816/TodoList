package models

import (
	"github.com/chuxin0816/web-tutorial/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func GetTodoList() (todoList []Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetTodo(id string) (todo Todo, err error) {
	err = dao.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	return
}

func PostTodo(todo Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func PutTodo(todo Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
