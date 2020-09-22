package dao

import "database/sql"

const (
	table = "albums"
)

type Album struct {
	Id    int    `db:"id" json:"id,omitempty"`
	Genre string `db:"genre" json:"genre,omitempty"`
	Name  string `db:"name" json:"name,omitempty"`
}

type AlbumDao struct{}

func (d AlbumDao) ListAll() []Album {
	albums := []Album{}
	err := dbCon.Select(&albums, "SELECT id,genre,name FROM "+table)
	if err != nil {
		err.Error()
	}

	return albums
}

func (d AlbumDao) GetById(id int) Album {
	album := Album{}
	err := dbCon.Get(&album, "SELECT id,genre,name FROM "+table+" WHERE id=?", id)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		err.Error()
	}

	return album
}
