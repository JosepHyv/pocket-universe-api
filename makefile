.PHONY: docs run

docs:
	swag init -g cmd/server/main.go --parseDependency --parseInternal

run: docs
	go run cmd/server/main.go
