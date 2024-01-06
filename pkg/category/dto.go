package category

type CreateCategoryDTO struct {
    Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateCategoryDTO struct {
    Name string `json:"name" validate:"min=3,max=100"`
}