.PHONE: run
run:
	go run cmd/gophermart/main.go

test:
	go test ./...

up:
	docker-compose up -d

down:
	docker-compose down