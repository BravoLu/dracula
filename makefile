protoc-gen:
	protoc -I=./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-grpc_opt=require_unimplemented_servers=false dracula.proto
