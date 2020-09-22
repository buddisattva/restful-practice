package dao

import (
	"database/sql"
)

const (
	table = "albums"
)

type Album struct {
	Id    int    `db:"id" json:"id,omitempty"`
	Genre string `db:"genre" json:"genre,omitempty"`
	Name  string `db:"name" json:"name,omitempty"`
}

type AlbumDao struct{}

func (d AlbumDao) Store(input map[string]interface{}) bool {
	_, err := dbCon.Query("INSERT INTO "+table+" (genre,name) VALUES (?,?)", input["Genre"], input["Name"])
	if err != nil {
		panic(err.Error())
	}

	return true
}

func (d AlbumDao) ListAll() []Album {
	albums := []Album{}
	err := dbCon.Select(&albums, "SELECT id,genre,name FROM "+table)
	if err != nil {
		panic(err.Error())
	}

	return albums
}

func (d AlbumDao) GetById(id int) Album {
	album := Album{}
	err := dbCon.Get(&album, "SELECT id,genre,name FROM "+table+" WHERE id=?", id)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		panic(err.Error())
	}

	return album
}

func (d AlbumDao) DeleteById(id int) bool {
	_, err := dbCon.Query("DELETE FROM "+table+" WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	return true
}
