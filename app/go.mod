module restful

go 1.15

replace controller => ./http/controller

replace gateway => ./adapter/gateway

replace dao => ./adapter/gateway/dao

replace input => ./input

require (
	controller v0.0.0-00010101000000-000000000000
	dao v0.0.0-00010101000000-000000000000 // indirect
	gateway v0.0.0-00010101000000-000000000000 // indirect
	github.com/zenazn/goji v1.0.1
	input v0.0.0-00010101000000-000000000000 // indirect
)
