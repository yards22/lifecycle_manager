sqlc:
	sqlc generate

dev:
	go run cmd/http-server/*.go

dev_windows:
	cd cmd\http-server\ && go run .

feb:
	cd web && yarn build && cp -r build ../build/

