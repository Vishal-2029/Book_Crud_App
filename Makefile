hello:
	eacho "hello"

build:
	go build -o BookCrud_App cmd/main.go

run:
	go run cmd/main.go

docker:
	docker compose up --build