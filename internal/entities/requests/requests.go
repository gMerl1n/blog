package requests

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type UpdatePostRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	TitleBody string `json:"title_body,omitempty"`
}
