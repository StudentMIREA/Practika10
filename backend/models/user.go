package models

type User struct {
	ID    int    `db:"id" json:"ID"`
	Image string `db:"image" json:"Image"`
	Name  string `db:"name" json:"Name"`
	Phone string `db:"phone" json:"Phone"`
	Mail  string `db:"mail" json:"Mail"`
}
