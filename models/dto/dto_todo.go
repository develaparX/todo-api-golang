package dto

type TodoRequest struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
}

type CreateTodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

type UpdateTodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
