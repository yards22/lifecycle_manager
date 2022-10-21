sqlc:
	sqlc generate

dev:
	go run cmd/http-server/*.go

feb:
	cd web && yarn build && cp -r build ../build/