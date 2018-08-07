package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTasks(pageNum int, pageSize int, maps interface{}) (tasks []Task) {
	db.Where("created_by = ?", maps).Find(&tasks)
	// db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tasks)
	return
}

func GetTaskTotal(maps interface{}) (count int) {
	db.Model(&Task{}).Where(maps).Count(&count)
	return
}

func ExistTaskByName(name string) bool {
	var task Task
	db.Select("id").Where("name = ?", name).First(&task)
	if task.ID > 0 {
		return true
	}

	return false
}

func AddTask(name string, state int, createdBy string) bool {
	db.Create(&Task{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func ExistTaskByID(id int) bool {
	var task Task
	db.Select("id").Where("id = ?", id).First(&task)
	if task.ID > 0 {
		return true
	}

	return false
}

func DeleteTask(id int) bool {
	db.Where("id = ?", id).Delete(&Task{})

	return true
}

func EditTask(id int, data interface{}) bool {
	db.Model(&Task{}).Where("id = ?", id).Updates(data)

	return true
}

func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())

	return nil
}

func (task *Task) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedAt", time.Now().Unix())
	return nil
}
