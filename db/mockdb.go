package db

// MockDB is a mock implementation of the Database interface for testing
type MockDB struct {
	CreateTaskFunc func(task string) (int, error)
	AllTasksFunc   func() ([]Task, error)
	DeleteTaskFunc func(key int) error
}

// CreateTask is a mock implementation of the CreateTask method
func (m *MockDB) CreateTask(task string) (int, error) {
	if m.CreateTaskFunc != nil {
		return m.CreateTaskFunc(task)
	}
	return 0, nil // Mock implementation, return default values
}

// AllTasks is a mock implementation of the AllTasks method
func (m *MockDB) AllTasks() ([]Task, error) {
	if m.AllTasksFunc != nil {
		return m.AllTasksFunc()
	}
	return nil, nil // Mock implementation, return default values
}

// DeleteTask is a mock implementation of the DeleteTask method
func (m *MockDB) DeleteTask(key int) error {
	if m.DeleteTaskFunc != nil {
		return m.DeleteTaskFunc(key)
	}
	return nil // Mock implementation, return default values
}

func GetMockDB() Database {
	return &MockDB{
		CreateTaskFunc: func(task string) (int, error) {
			return 1, nil
		},
		AllTasksFunc: func() ([]Task, error) {
			tasks := []Task{
				{Id: 1, Value: "Task 1"},
				{Id: 2, Value: "Task 2"},
			}
			return tasks, nil
		},
		DeleteTaskFunc: func(key int) error {
			return nil
		},
	}
}
