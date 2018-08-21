package todo_service

import "Gin-Todo/models"

type Todo struct {
	Id        string
	UserID    string
	Name      string
	Desc      string
	Completed bool

	PageNum  int
	PageSize int
}

func (t *Todo) Create() error {
	todo := map[string]interface{}{
		"user_id":   t.UserID,
		"name":      t.Name,
		"desc":      t.Desc,
		"completed": t.Completed,
	}
	if err := models.AddTodo(todo); err != nil {
		return err
	}
	return nil
}

func (t *Todo) Update() error {
	return models.EditTodo(t.Id, map[string]interface{}{
		"id":        t.Id,
		"name":      t.Name,
		"desc":      t.Desc,
		"completed": t.Completed,
	})
}

func (t *Todo) Find() (*models.Todo, error) {
	todo, err := models.GetTodo(t.Id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Todo) FindAll() ([]*models.Todo, error) {
	todo, err := models.GetTodos(t.PageNum, t.PageSize, map[string]interface{}{"user_id": t.UserID})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *Todo) Delete() error {
	return models.DeleteTodo(t.Id)
}

func (t *Todo) Count() (int, error) {
	return models.GetTodoTotal(map[string]interface{}{"user_id": t.UserID})
}

func (a *Todo) ExistByID() (bool, error) {
	return models.ExistByID(a.Id)
}
