sqlc:
	sqlc generate

dev:
	go run cmd/http-server/*.go

dev_windows:
	cd cmd\http-server\ && go run .

lcm:
	rm -rf app
	mkdir app
	make beb

beb:
	go build -o app/api cmd/http-server/*.go
