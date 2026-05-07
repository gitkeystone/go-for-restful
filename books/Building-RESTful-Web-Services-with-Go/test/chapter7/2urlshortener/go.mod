module github.com/gitkeystone/url-shortener

go 1.26.1

replace helper => ./helper

replace utils => ./utils

require (
	github.com/gorilla/mux v1.8.1
	helper v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

require github.com/lib/pq v1.12.3 // indirect
