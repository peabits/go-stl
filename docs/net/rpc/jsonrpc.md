
# jsonrpc

## 概述

jsonrpc 包为 rpc 包实现了 JSON-RPC 1.0 ClientCodec 和 ServerCodec。有关 JSON-RPC 2.0 的支持，请参阅 https://godoc.org/?q=json-rpc+2.0

## 索引

- [func Dial(network, address string) (*rpc.Client, error)](#func-dial)
- [func NewClient(conn io.ReadWriteCloser) *rpc.Client](#func-newclient)
- [func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec](#func-newclientcodec)
- [func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec](#func-newservercodec)
- [func ServeConn(conn io.ReadWriteCloser)](#func-serveconn)

## 常量

## 变量

## 函数

### func Dial

```go
func Dial(network, address string) (*rpc.Client, error)
```

Dial 连接到位于指定网络地址的 JSON-RPC 服务器。

### func NewClient

```go
func NewClient(conn io.ReadWriteCloser) *rpc.Client
```

NewClient 返回一个新的 rpc.Client 来处理对连接另一端的服务集的请求。

### func NewClientCodec

```go
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec
```

NewClientCodec 在 conn 上使用 JSON-RPC 返回一个新的 rpc.ClientCodec。

### func NewServerCodec

```go
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec
```

NewServerCodec 在 conn 上使用 JSON-RPC 返回一个新的 rpc.ServerCodec。

### func ServeConn

```go
func ServeConn(conn io.ReadWriteCloser)
```

ServeConn 在单个连接上运行 JSON-RPC 服务器。 ServeConn 阻塞，服务连接直到客户端挂断。调用者通常在 go 语句中调用 ServeConn。

## 类型
