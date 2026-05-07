module github.com/gitkeystone/basic-example

go 1.26.1

replace helper => ./helper

require helper v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.12.3 // indirect
