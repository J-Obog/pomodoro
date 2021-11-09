package task

type Task struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}