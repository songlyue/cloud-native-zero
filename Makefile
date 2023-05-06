go:
	goctl api go --api ./app/testApi/cmd/api/desc/test.api  -dir ./app/testApi/cmd/api
go12:
	goctl rpc protoc desc/test.proto --go_out=./client --go-grpc_out=./client -zrpc_out=./client

model:
	goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/test_as" -table="user" -dir ./app/testApi/cmd/api/model
