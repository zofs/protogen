.PHONY: dtopb 
dtopb:
	protoc --proto_path=protos -Iproto protos/dtopb/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative 

.PHONY: accore 
accore:
	protoc --proto_path=protos -Iproto protos/accore/**/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative --grpc-gateway_out .  --grpc-gateway_opt paths=source_relative --openapiv2_out=./dist --openapiv2_opt allow_merge=true,merge_file_name=accore

.PHONY: todo 
todo:
	protoc --proto_path=protos -Iproto protos/todo/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative --grpc-gateway_out .  --grpc-gateway_opt paths=source_relative --openapiv2_out=./dist --openapiv2_opt allow_merge=true,merge_file_name=todo

.PHONY: all
all: dtopb accore