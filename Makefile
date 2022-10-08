sqlc:
	sqlc generate

dev:
	cd cmd/http-server/ && go run *.go
