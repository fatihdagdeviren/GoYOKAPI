We will use Go modules to handle this dependency. Running the following commands will create the go.mod and go.sum files:
go mod init GoYOKAPI
go mod tidy

docker build -t goyokapi .
docker run -t -p 8089:8000 goyokapi

--fatih