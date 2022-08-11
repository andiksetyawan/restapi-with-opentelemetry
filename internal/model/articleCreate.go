package model

type ArticleCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
