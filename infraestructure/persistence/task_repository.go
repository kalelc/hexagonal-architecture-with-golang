package persistence

import "github.com/kalelc/hexagonal-arquitecture/domain"

type TaskRepository struct {
	task []*domain.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) Save(task *domain.Task) error {
	r.task = append(r.task, task)
	return nil
}

func (r *TaskRepository) FindAll() ([]*domain.Task, error) {
	return r.task, nil
}
