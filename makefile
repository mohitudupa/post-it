setup:
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/gin-gonic/gin
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
dev:
	go run notes/httpd
swagger:
	swagger generate spec -o ./swagger.yml
build:
	env GOARCH="amd64" GOOS="darwin" go build -o builds/notes-mac postit
	env GOARCH="amd64" GOOS="windows" go build -o builds/notes-windows postit
	env GOARCH="amd64" GOOS="linux" go build -o builds/notes-linux postit
