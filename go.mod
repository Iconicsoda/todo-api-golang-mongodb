module github.com/iconicsoda/todo-api-golang-mongodb

go 1.15

require github.com/google/uuid v1.2.0 // indirect
require Routes/Todo v1.0.0
replace Routes/Todo => ./Routes/Todo
