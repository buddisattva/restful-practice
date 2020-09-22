package dao

const (
	table = "albums"
)

type Album struct {
	Id    int    `db:"id"`
	Genre string `db:"genre"`
	Name  string `db:"name"`
}

type AlbumDao struct{}

func (d AlbumDao) ListAll() []Album {
	albums := []Album{}
	err := dbCon.Select(&albums, "SELECT id,genre,name FROM "+table)
	if err != nil {
		panic(err.Error())
	}

	return albums
}
