# protoc command:

```
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_opt=module=github.com/Clement-Jean/grpc-go-course --go-grpc_out=. greet/proto/dummy.proto
```