unit-test:
	go test -v -coverprofile=coverage.out ./...

run-docker:
	docker-compose up -d

run:
	go run cmd/main.go