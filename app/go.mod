module restful

go 1.15

replace controllers => ./http/controllers

require (
	controllers v0.0.0-00010101000000-000000000000
	github.com/zenazn/goji v1.0.1
)
