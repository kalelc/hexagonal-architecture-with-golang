package service

import (
	"github.com/kalelc/hexagonal-arquitecture/application/command"
	"github.com/kalelc/hexagonal-arquitecture/application/query"
	"github.com/kalelc/hexagonal-arquitecture/domain"
)

type TaskService struct {
	repository domain.TaskRepository
}

func NewTaskService(repository domain.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) HandleCreateTask(cmd command.CreateTaskCommand) error {
	task := domain.NewTask(cmd.ID, cmd.Title, cmd.Description)
	return s.repository.Save(task)
}

func (s *TaskService) HandleGetTasks(q query.GetTasksQuery) ([]*domain.Task, error) {
	return s.repository.FindAll()
}
