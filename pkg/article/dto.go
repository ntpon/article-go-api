package article

type CreateArticleDTO struct {
    CategoryID int    `json:"category_id" validate:"required"`
    Title      string `json:"title" validate:"required,min=3,max=200"`
    Content    string `json:"content" validate:"required,min=10"`
}

type UpdateArticleDTO struct {
    CategoryID int    `json:"category_id" validate:"required"`
    Title      string `json:"title" validate:"min=3,max=200"`
    Content    string `json:"content" validate:"min=10"`
}