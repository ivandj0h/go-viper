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

**Generate Proto**

```go
protoc --proto_path=protos protos/*.proto --go_out=.
```

**Run The App**

```go
go run main.go
```
