gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative companypb/company.proto

remove:
	rm companypb/*.go	