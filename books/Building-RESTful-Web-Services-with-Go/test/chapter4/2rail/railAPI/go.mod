module github.com/gitkeystone/rail-api

go 1.26.1

replace dbutils => ../dbutils

require (
	dbutils v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.42
)

require github.com/emicklei/go-restful/v3 v3.13.0
