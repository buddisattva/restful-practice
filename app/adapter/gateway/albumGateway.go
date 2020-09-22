package gateway

import "dao"

type AlbumGateway struct{}

func (g AlbumGateway) ListAll() []dao.Album {
	albumDao := new(dao.AlbumDao)

	return albumDao.ListAll()
}
