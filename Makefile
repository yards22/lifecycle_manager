sqlc:
	sqlc generate

dev:
	go run cmd/http-server/*.go

dev_windows:
	cd cmd\http-server\ && go run .

lcm:
	rm -rf app
	mkdir app
	make feb
	make beb

feb:
	cd web && yarn && yarn build && cp -r build ../app/

beb:
	go build -o app/api cmd/http-server/*.go
