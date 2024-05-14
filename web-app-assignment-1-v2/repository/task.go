package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	t.filebased.UpdateTask(task.ID, *task)
	return nil
}

func (t *taskRepository) Delete(id int) error {
	t.filebased.DeleteTask(id)
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil // TODO: replace this
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks, _ := t.filebased.GetTasks()
	return tasks, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	category, _ := t.filebased.GetTaskListByCategory(id)

	return category, nil
}
