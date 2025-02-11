APP_SERVER_NAME=godb-server
APP_CLIENT_NAME=godb-client

build-server:
	go build -o ${APP_SERVER_NAME} cmd/db/server/main.go

build-client:
	go build -o ${APP_CLIENT_NAME} cmd/db/client/main.go