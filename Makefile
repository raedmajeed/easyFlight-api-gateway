run:
	cd cmd/api && go run main.go

proto admin:
	cd pkg/admin/pb && protoc --go_out=. --go-grpc_out=. admin_airline.proto

protoc booking:
	cd pkg/bookingService/pb && protoc --go-grpc_out=. --go_out=. booking.proto

adminSwag:
	swag init -g ../cmd/api/main.go -o ./docs -d ./pkg

bookingSwag:
	swag init -g ../../cmd/api/main.go -o ./docs -d ./pkg/bookingService -generalInfo ./booking.json
