PROTOC_GEN_GO="/Users/tanasinp/go/bin/protoc-gen-go"
PROTOC_GEN_GO_GRPC="/Users/tanasinp/go/bin/protoc-gen-go-grpc"

protoc \
  --plugin=protoc-gen-go=$PROTOC_GEN_GO \
  --plugin=protoc-gen-go-grpc=$PROTOC_GEN_GO_GRPC \
  --proto_path=payment \
  --go_out=payment --go_opt=paths=source_relative \
  --go-grpc_out=payment --go-grpc_opt=paths=source_relative \
  payment/payment.proto

protoc \
  --plugin=protoc-gen-go=$PROTOC_GEN_GO \
  --plugin=protoc-gen-go-grpc=$PROTOC_GEN_GO_GRPC \
  --proto_path=order \
  --go_out=order --go_opt=paths=source_relative \
  --go-grpc_out=order --go-grpc_opt=paths=source_relative \
  order/order.proto
