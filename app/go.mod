module restful

go 1.15

replace controller => ./http/controller
replace gateway => ./adapter/gateway
replace dao => ./adapter/gateway/dao

require (
	controller v0.0.0-00010101000000-000000000000
	github.com/zenazn/goji v1.0.1
)
