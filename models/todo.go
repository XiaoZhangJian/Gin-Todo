package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	// Model
	Id        string `json:"id"`
	UserID    string `json:"user_id" gorm:"index"`
	Name      string `json:"name"`      // 标题
	Desc      string `json:"desc"`      // 描述
	Completed bool   `json:"completed"` // 状态

}

func AddTodo(data map[string]interface{}) error {
	todo := Todo{
		Id:     NewId(),
		UserID: data["user_id"].(string),
		Name:   data["name"].(string),
		Desc:   data["desc"].(string),
	}

	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(id string) (*Todo, error) {
	var todo Todo
	err := db.Where("id = ? ", id).First(&todo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &todo, nil
}

func GetTodos(pageNum int, pageSize int, maps interface{}) ([]*Todo, error) {
	var todos []*Todo
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&todos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return todos, nil
}

func GetTodoTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Todo{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ExistByID(id string) (bool, error) {
	var todo Todo
	err := db.Select("id").Where("id = ?", id).First(&todo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if todo.Id == id {
		return true, nil
	}

	return false, nil
}

func DeleteTodo(id string) error {
	if err := db.Where("id = ?", id).Delete(&Todo{}).Error; err != nil {
		return err
	}
	return nil
}

func EditTodo(id string, data interface{}) error {
	if err := db.Model(&Todo{}).Where("id = ?  ", id).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func NewId() string {
	id, _ := uuid.NewV4()
	return id.String()
}
