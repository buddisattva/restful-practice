package gateway

import "dao"

var albumDao *dao.AlbumDao

func init() {
	albumDao = new(dao.AlbumDao)
}

type AlbumGateway struct{}

func (g AlbumGateway) ListAll() []dao.Album {
	return albumDao.ListAll()
}

func (g AlbumGateway) GetById(id int) dao.Album {
	return albumDao.GetById(id)
}

func (g AlbumGateway) DeleteById(id int) bool {
	return albumDao.DeleteById(id)
}
