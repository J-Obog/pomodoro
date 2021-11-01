package tasksvc

type Task struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}