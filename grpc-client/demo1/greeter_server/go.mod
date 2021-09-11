module github.com/bigwhite/grpc-client/demo1

go 1.17

require (
	google.golang.org/grpc v1.40.0
	google.golang.org/grpc/examples v1.40.0
)

require (
	github.com/golang/protobuf v1.4.3 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace google.golang.org/grpc v1.40.0 => /Users/tonybai/Go/src/github.com/grpc/grpc-go

replace google.golang.org/grpc/examples v1.40.0 => /Users/tonybai/Go/src/github.com/grpc/grpc-go/examples
