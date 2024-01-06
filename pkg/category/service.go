package category

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
    db *sqlx.DB
}

// NewService creates a new category service
func NewService(db *sqlx.DB) *Service {
    return &Service{db: db}
}

// CreateCategory adds a new category to the database
func (s *Service) CreateCategory(cat *Category) error {
    _, err := s.db.NamedExec(`INSERT INTO categories (name) VALUES (:name)`, cat)
    return err
}

// GetAllCategories retrieves all categories from the database
func (s *Service) GetAllCategories() ([]Category, error) {
    var categories []Category
    err := s.db.Select(&categories, "SELECT * FROM categories")
    return categories, err
}

// GetCategory retrieves a single category by ID
func (s *Service) GetCategory(id int) (Category, error) {
    var cat Category
    err := s.db.Get(&cat, "SELECT * FROM categories WHERE id = ?", id)    
    return cat, err
}

// UpdateCategory updates an existing category
func (s *Service) UpdateCategory(cat *Category) error {
    _, err := s.db.NamedExec(`UPDATE categories SET name = :name WHERE id = :id`, cat)
    return err
}

// DeleteCategory deletes a category by ID
func (s *Service) DeleteCategory(id int) error {
    _, err := s.db.Exec("DELETE FROM categories WHERE id = ?", id)
    return err
}
