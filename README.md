## Go App with GRPC

### About

This is Simple CRUD App using `Go` with `GRPC`

The Steps are :

**App Init**

```go
go mod init
```

```go
go mod tidy
```

**Compailer Plugin**

```go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

**Generate Proto**

```go
protoc --proto_path=protos protos/*.proto --go_out=.
```

**Run The App**

```go
go run main.go
```

```go
Masih terus dalam proses Pengembangan...
```
