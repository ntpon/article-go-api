package category

type Category struct {
    ID   int    `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
}
