package gateway

import dao

type albumGateway struct{}

func init() {
	dao = dao.albumDao
}

func (g albumGateway) ListAll() string {
	return albumDao.ListAll()
}
