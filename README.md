# protoc command:

```
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/rafaelmbcosta/grpc-go-studiese --go-grpc_opt=module=github.com/rafaelmbcosta/grpc-go-studies --go-grpc_out=. greet/proto/dummy.proto
```


remember to run "go mod tidy"