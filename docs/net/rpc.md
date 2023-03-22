
# rpc

## 概述

rpc 包提供通过网络或其他 I/O 连接访问对象的导出方法。服务器注册一个对象，使其作为具有对象类型名称的服务可见。注册后，对象的导出方法将可以远程访问。服务器可以注册多个不同类型的对象（服务），但注册多个相同类型的对象是错误的。

只有满足这些标准的方法才能用于远程访问；其他方法将被忽略：

- 方法的类型被导出。
- 该方法已导出。
- 该方法有两个参数，都是导出（或内置）类型。
- 该方法的第二个参数是一个指针。
- 该方法有返回类型错误。

实际上，该方法必须看起来像

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

其中 T1 和 T2 可以通过 encoding/gob 进行编组。即使使用不同的编解码器，这些要求也适用。 （将来，这些要求可能会针对自定义编解码器放宽。）

其中 T1 和 T2 可以通过 encoding/gob 进行编组。即使使用不同的编解码器，这些要求也适用。 （将来，这些要求可能会针对自定义编解码器放宽。）

方法的第一个参数表示调用者提供的参数；第二个参数表示返回给调用者的结果参数。该方法的返回值，如果非零，则作为客户端看到的字符串传回，就像由 errors.New 创建的一样。如果返回错误，则不会将 reply 参数发送回客户端。

服务器可以通过调用 ServeConn 处理单个连接上的请求。更典型的是，它将创建一个网络侦听器并调用 Accept，或者对于 HTTP 侦听器，HandleHTTP 和 http.Serve。

希望使用该服务的客户端建立连接，然后在连接上调用 NewClient。便捷函数 Dial (DialHTTP) 为原始网络连接（HTTP 连接）执行这两个步骤。生成的 Client 对象有两个方法，Call 和 Go，它们指定要调用的服务和方法，一个包含参数的指针，以及一个接收结果参数的指针。

Call 方法等待远程调用完成，而 Go 方法异步启动调用并使用 Call 结构的 Done 通道发出完成信号。

除非设置了明确的编解码器，否则使用包 encoding/gob 来传输数据。

这是一个简单的例子。服务器希望导出 Arith 类型的对象：

```go
package server

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
```

服务器调用（对于 HTTP 服务）：

```go
arith := new(Arith)
rpc.Register(arith)
rpc.HandleHTTP()
l, e := net.Listen("tcp", ":1234")
if e != nil {
	log.Fatal("listen error:", e)
}
go http.Serve(l, nil)
```

此时，客户端可以看到服务 “Arith”，其方法为 “Arith.Multiply” 和 “Arith.Divide”。要调用一个，客户端首先拨打服务器：

```go
client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
if err != nil {
	log.Fatal("dialing:", err)
}
```

然后它可以进行远程调用：

```go
// Synchronous call
args := &server.Args{7,8}
var reply int
err = client.Call("Arith.Multiply", args, &reply)
if err != nil {
	log.Fatal("arith error:", err)
}
fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
```

或

```go
// Asynchronous call
quotient := new(Quotient)
divCall := client.Go("Arith.Divide", args, quotient, nil)
replyCall := <-divCall.Done	// will be equal to divCall
// check errors, print, etc.
```

服务器实现通常会为客户端提供一个简单的、类型安全的包装器。

net/rpc 包已冻结，不接受新功能。

## 索引

- [Constants](#常量)
- [Variables](#变量)
- [Functions](#函数)
    - [func Accept(lis net.Listener)]()
    - [func HandleHTTP()]()
    - [func Register(rcvr any) error]()
    - [func RegisterName(name string, rcvr any) error]()
    - [func ServeCodec(codec ServerCodec)]()
    - [func ServeConn(conn io.ReadWriteCloser)]()
    - [func ServeRequest(codec ServerCodec) error]()
- [Types](#类型)
    - [type Call]()
    - [type Client]()
        - [func Dial(network, address string) (*Client, error)]()
        - [func DialHTTP(network, address string) (*Client, error)]()
        - [func DialHTTPPath(network, address, path string) (*Client, error)]()
        - [func NewClient(conn io.ReadWriteCloser) *Client]()
        - [func NewClientWithCodec(codec ClientCodec) *Client]()
        - [func (client *Client) Call(serviceMethod string, args any, reply any) error]()
        - [func (client *Client) Close() error]()
        - [func (client *Client) Go(serviceMethod string, args any, reply any, done chan *Call) *Call]()
    - [type ClientCodec]()
    - [type Request]()
    - [type Response]()
    - [type Server]()
        - [func NewServer() *Server]()
        - [func (server *Server) Accept(lis net.Listener)]()
        - [func (server *Server) HandleHTTP(rpcPath, debugPath string)]()
        - [func (server *Server) Register(rcvr any) error]()
        - [func (server *Server) RegisterName(name string, rcvr any) error]()
        - [func (server *Server) ServeCodec(codec ServerCodec)]()
        - [func (server *Server) ServeConn(conn io.ReadWriteCloser)]()
        - [func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)]()
        - [func (server *Server) ServeRequest(codec ServerCodec) error]()
    - [type ServerCodec]()
    - [type ServerError]()
        - [func (e ServerError) Error() string]()

## 常量

```go
const (
	// Defaults used by HandleHTTP
	DefaultRPCPath   = "/_goRPC_"
	DefaultDebugPath = "/debug/rpc"
)
```

## 变量

```go
var DefaultServer = NewServer()
```

DefaultServer 是 *Server 的默认实例。

```go
var ErrShutdown = errors.New("connection is shut down")
```

## 函数

### func Accept

### func HandleHTTP

### func Register

### func RegisterName

### func ServeCodec

### func ServeConn

### func ServeRequest

## 类型

### type Call

### type Client

### func Dial

### func DialHTTP

### func DialHTTPPath

### func NewClient

### func NewClientWithCodec

### func (*Client) Call

### func (*Client) Close

### func (*Client) Go

### type ClientCodec

### type Request

### type Response

### type Server

### func NewServer

### func (*Server) Accept

### func (*Server) HandleHTTP

### func (*Server) Register

### func (*Server) RegisterName

### func (*Server) ServeCodec

### func (*Server) ServeConn

### func (*Server) ServeHTTP

### func (*Server) ServeRequest

### type ServerCodec

### type ServerError

### func (ServerError) Error

## 目录

- [jsonrpc](rpc/jsonrpc.md)
