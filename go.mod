module github.com/Pogodoc/pogodoc-go

go 1.24

require (
	github.com/Pogodoc/pogodoc-go/client v0.3.0
	github.com/joho/godotenv v1.5.1
)

replace github.com/Pogodoc/pogodoc-go/client => ./client

require github.com/google/uuid v1.4.0 // indirect
