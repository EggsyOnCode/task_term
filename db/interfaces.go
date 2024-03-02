package db

type Database interface {
	CreateTask(task string) (int, error)
	AllTasks() ([]Task, error)
	DeleteTask(key int) error
}

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}
