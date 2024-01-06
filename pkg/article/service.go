package article

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
    db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
    return &Service{db: db}
}

func (s *Service) CreateArticle(a *Article) error {
    _, err := s.db.NamedExec(`INSERT INTO articles (category_id, title, content) VALUES (:category_id, :title, :content)`, a)
    return err
}

func (s *Service) GetAllArticles() ([]Article, error) {
    var articles []Article
    err := s.db.Select(&articles, "SELECT * FROM articles")
    return articles, err
}

func (s *Service) GetArticle(id int) (Article, error) {
    var article Article
    err := s.db.Get(&article, "SELECT * FROM articles WHERE id = ?", id)
    return article, err
}

func (s *Service) UpdateArticle(a *Article) error {
    _, err := s.db.NamedExec(`UPDATE articles SET category_id = :category_id, title = :title, content = :content WHERE id = :id`, a)
    return err
}

func (s *Service) DeleteArticle(id int) error {
    _, err := s.db.Exec("DELETE FROM articles WHERE id = ?", id)
    return err
}
