package domain

type Task struct {
	ID          string
	Title       string
	Description string
}

func NewTask(id, title, description string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
	}
}
