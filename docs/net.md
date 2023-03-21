
# net

## 概述

net 包为网络 I/O 提供了一个可移植的接口，包括 TCP/IP、UDP、域名解析和 Unix 域套接字。

尽管该包提供对低级网络原语的访问，但大多数客户端只需要 Dial、Listen 和 Accept 函数以及关联的 Conn 和 Listener 接口提供的基本接口。 crypto/tls 包使用相同的接口和类似的 Dial 和 Listen 函数。

Dial 函数连接到服务器：

```go
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
    // handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
```

Listen 函数创建服务器：

```go
ln, err := net.Listen("tcp", ":8080")
if err != nil {
    // handle error
}
for {
    conn, err := ln.Accept()
    if err != nil {
        // handle error
    }
    go handleConnection()
}
```

### 域名解析

解析域名的方法，是间接使用 Dial 等函数，还是直接使用 LookupHost、LookupAddr 等函数，因操作系统而异。

在 Unix 系统上，解析器有两个解析名称的选项。它可以使用纯 Go 解析器将 DNS 请求直接发送到 /etc/resolv.conf 中列出的服务器，也可以使用基于 cgo 的解析器调用 C 库例程，例如 getaddrinfo 和 getnameinfo。

默认情况下使用纯 Go 解析器，因为阻塞的 DNS 请求只消耗一个 goroutine，而阻塞的 C 调用消耗一个操作系统线程。当 cgo 可用时，基于 cgo 的解析器将在多种情况下使用：在不允许程序直接发出 DNS 请求的系统上 (OS X)，当存在 LOCALDOMAIN 环境变量时（即使是空的），当RES_OPTIONS 或 HOSTALIASES 环境变量非空，当 ASR_CONFIG 环境变量非空时（仅限 OpenBSD），当 /etc/resolv.conf 或 /etc/nsswitch.conf 指定使用 Go 解析器未实现的功能时, 并且当被查找的名称以 .local 结尾或者是一个 mDNS 名称时。

可以通过将 GODEBUG 环境变量（参见包运行时）的 netdns 值设置为 go 或 cgo 来覆盖解析器决策，如下所示：

```go
export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force native resolver (cgo, win32)
```

通过设置 netgo 或 netcgo 构建标记，在构建 Go 源代码树时也可以强制执行该决定。

数字 netdns 设置，如 GODEBUG=netdns=1，会导致解析器打印有关其决策的调试信息。要在打印调试信息的同时强制使用特定的解析器，请通过加号连接两个设置，如 GODEBUG=netdns=go+1。

在 Plan 9 上，解析器总是访问 /net/cs 和 /net/dns。

在 Windows 上，在 Go 1.18.x 及更早版本中，解析器始终使用 C 库函数，例如 GetAddrInfo 和 DnsQuery。

## 索引

- [Contants](#constants)
- [Variables](#variables)
- [Functions](#functions)
    - [func JoinHostPort(host, port string) string]()
    - [func LookupAddr(addr string) (names []string, err error)]()
    - [func LookupCNAME(host string) (cname string, err error)]()
    - [func LookupHost(host string) (addrs []string, err error)]()
    - [func LookupPort(network, service string) (port int, err error)]()
    - [func LookupTXT(name string) ([]string, error)]()
    - [func ParseCIDR(s string) (IP, *IPNet, error)]()
    - [func Pipe() (Conn, Conn)]()
    - [func SplitHostPort(hostport string) (host, port string, err error)]()
- [Types](#types) 
    - [type Addr]()
        - [func InterfaceAddrs()([]Addr, error)]()
    - [type AddrError]()
## Constants

IP 地址长度（字节）。

```go
const (
    IPv4len = 4
    IPv6len = 16
)
```

## Variables

众所周知的 IPv4 地址

```go
var (
	IPv4bcast     = IPv4(255, 255, 255, 255) // limited broadcast
	IPv4allsys    = IPv4(224, 0, 0, 1)       // all systems
	IPv4allrouter = IPv4(224, 0, 0, 2)       // all routers
	IPv4zero      = IPv4(0, 0, 0, 0)         // all zeros
)
```

众所周知的 IPv6 地址

```go
var (
	IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)
```

DefaultResolver 是包级查找函数和没有指定解析器的拨号器使用的解析器。

```go
var DefaultResolver = &Resolver{}
```

ErrClosed 是 I/O 调用在已经关闭的网络连接上返回的错误，或者在 I/O 完成之前被另一个 goroutine 关闭。这可能包含在另一个错误中，通常应该使用 errors.Is(err, net.ErrClosed) 进行测试。

```go
var ErrClosed error = errClosed
```

OpError 中包含的各种错误。

```go
var (
	ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
)
```

## Functions

## Types

### type Addr

```go
type Addr interface {
	Network() string // name of the network (for example, "tcp", "udp")
	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}
```

Addr 表示网络端点地址。

Network 和 String 这两个方法通常返回可以作为参数传递给 Dial 的字符串，但字符串的确切形式和含义取决于实现。

### func InterfaceAddrs

```go
func InterfaceAddrs() ([]Addr, error)
```

InterfaceAddrs 返回系统的单播接口地址列表。

返回的列表没有标识关联的接口；使用 Interfaces 和 Interface.Addrs 了解更多细节。

### type AddrError

```go
type AddrError struct {
	Err  string
	Addr string
}
```

### func (*AddrError) Error

### func (*AddrError) Temporary

### func (*AddrError) Timeout

### type Buffers

### func (*Buffers) Read

### func (*Buffers) WriteTo


