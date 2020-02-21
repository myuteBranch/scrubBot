package utils

import (
	"github.com/jmoiron/sqlx/types"
)

// LinkObj for db model
type LinkObj struct {
	ID        int
	Username  string
	Link      string
	Descrip   string
	CreatedOn string `db:"created_on"`
}

// ContactFavorites is a struct
type ContactFavorites struct {
	Colors []string `json:"colors"`
}

// Contact is a struct
type Contact struct {
	ID                   int
	Name, Address, Phone string

	FavoritesJSON types.JSONText    `db:"favorites"`
	Favorites     *ContactFavorites `db:"-"`

	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
