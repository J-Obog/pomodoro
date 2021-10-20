package api

type Task struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	//Description string `json:"description"` removing description for now
	Completed bool `json:"completed"`
}