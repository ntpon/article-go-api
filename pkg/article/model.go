package article

type Article struct {
    ID         int    `db:"id" json:"id"`
    CategoryID int    `db:"category_id" json:"category_id"`
    Title      string `db:"title" json:"title"`
    Content    string `db:"content" json:"content"`
}