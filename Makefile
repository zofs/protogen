.PHONY: dtopb 
dtopb:
	protoc --proto_path=protos -Iproto protos/dtopb/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative 

.PHONY: accore 
accore:
	protoc --proto_path=protos -Iproto protos/accore/**/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative --grpc-gateway_out .  --grpc-gateway_opt paths=source_relative --openapiv2_out=./dist --openapiv2_opt allow_merge=true,merge_file_name=accore

.PHONY: catalog 
catalog:
	protoc --proto_path=protos -Iproto protos/catalog/**/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative --grpc-gateway_out .  --grpc-gateway_opt paths=source_relative --openapiv2_out=./dist --openapiv2_opt allow_merge=true,merge_file_name=catalog

.PHONY: cart 
cart:
	protoc --proto_path=protos -Iproto protos/cart/**/*.proto  --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative --grpc-gateway_out .  --grpc-gateway_opt paths=source_relative --openapiv2_out=./dist --openapiv2_opt allow_merge=true,merge_file_name=cart

.PHONY: all
all: dtopb accore catalog cart