package entity

type Solution struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Tags  string `db:"tags" json:"tags"`
	Web   string `db:"web" json:"web"`
}
