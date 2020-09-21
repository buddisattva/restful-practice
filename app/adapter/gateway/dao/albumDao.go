package dao

import "database/sql"

type albumDao struct{}

func init() {
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())
	}
}

func (g albumGateway) ListAll() string {
	return "sw"
}
