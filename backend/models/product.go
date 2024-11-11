package models

type Product struct {
	ID          int     `db:"id" json:"ID"`
	Name        string  `db:"name" json:"Name"`
	Image       string  `db:"image" json:"Image"`
	Cost        float64 `db:"cost" json:"Cost"`
	Describtion string  `db:"description" json:"Describtion"`
	Favorite    bool    `db:"favorite" json:"Favorite"`
	ShopCart    bool    `db:"shop_cart" json:"ShopCart"`
	Count       int     `db:"count" json:"Count"`
}
