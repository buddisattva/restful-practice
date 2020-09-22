package gateway

import (
	"dao"
	"input"

	"github.com/fatih/structs"
)

var albumDao *dao.AlbumDao

func init() {
	albumDao = new(dao.AlbumDao)
}

type AlbumGateway struct{}

func (g AlbumGateway) Store(input input.NewAlbum) bool {
	mapInput := structs.Map(input)

	return albumDao.Store(mapInput)
}

func (g AlbumGateway) ListAll() []dao.Album {
	return albumDao.ListAll()
}

func (g AlbumGateway) GetById(id int) dao.Album {
	return albumDao.GetById(id)
}

func (g AlbumGateway) DeleteById(id int) bool {
	return albumDao.DeleteById(id)
}

func (g AlbumGateway) UpdateById(id int, input input.NewAlbum) bool {
	mapInput := structs.Map(input)

	return albumDao.UpdateById(id, mapInput)
}
