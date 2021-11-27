from golang:1.17.3-alpine3.14

run mkdir /app
workdir /app
run go mod init api_go
run go get -u github.com/gorilla/mux
cmd ["go mod tidy"]
copy . .
run go build app.go
cmd ["./app"]