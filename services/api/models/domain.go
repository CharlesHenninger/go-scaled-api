package models

type Domain struct {
	Id      int64  `db:"id"`
	Name    string `db:"name"`
	Events  int64  `db:"events"`
	Bounced bool   `db:"bounced"`
}
