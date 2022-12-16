sqlc:
	sqlc generate

dev:
<<<<<<< HEAD
	cd cmd\http-server\ && go run .
=======
	go run cmd/http-server/*.go

feb:
	cd web && yarn build && cp -r build ../build/
>>>>>>> bb664c37c41056be27acf7e6ef4a9b8dd2e55a57
