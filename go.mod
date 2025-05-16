module pogodoc/go

go 1.24

require (
	github.com/joho/godotenv v1.5.1
	pogodoc/go/sdk v0.0.0
)

require github.com/google/uuid v1.4.0 // indirect

replace pogodoc/go/sdk => ./sdk
