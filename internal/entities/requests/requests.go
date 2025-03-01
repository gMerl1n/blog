package requests

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type UpdatePostRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type CreateUserRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokensRefreshRequest struct {
	RefreshToken string
}
