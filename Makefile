run:
	cd cmd/api && go run main.go

proto admin:
	cd pkg/admin/pb && protoc --go_out=. --go-grpc_out=. admin_airline.proto