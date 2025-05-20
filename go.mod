module github.com/pogodoc/pogodoc-go

go 1.24

require (
	github.com/joho/godotenv v1.5.1
	github.com/pogodoc/pogodoc-go/sdk v0.0.0
)

replace github.com/pogodoc/pogodoc-go/sdk => ./sdk

require github.com/google/uuid v1.4.0 // indirect
