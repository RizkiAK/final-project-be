package domain

type LoginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateBlogRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
