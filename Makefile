all:
	go run cmd/main/main.go
	
migrate:
	go generate ./internal/storage/ent
	go run cmd/migrator/main.go