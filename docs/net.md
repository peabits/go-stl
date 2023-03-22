
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

```bash
export GODEBUG=netdns=go    # force pure Go resolver
export GODEBUG=netdns=cgo   # force native resolver (cgo, win32)
```

通过设置 netgo 或 netcgo 构建标记，在构建 Go 源代码树时也可以强制执行该决定。

数字 netdns 设置，如 GODEBUG=netdns=1，会导致解析器打印有关其决策的调试信息。要在打印调试信息的同时强制使用特定的解析器，请通过加号连接两个设置，如 GODEBUG=netdns=go+1。

在 Plan 9 上，解析器总是访问 /net/cs 和 /net/dns。

在 Windows 上，在 Go 1.18.x 及更早版本中，解析器始终使用 C 库函数，例如 GetAddrInfo 和 DnsQuery。

## 索引

- [Contants](#常量)
- [Variables](#变量)
- [Functions](#函数)
    - [func JoinHostPort(host, port string) string](#func-joinhostport)
    - [func LookupAddr(addr string) (names []string, err error)](#func-lookupaddr)
    - [func LookupCNAME(host string) (cname string, err error)](#func-lookupcname)
    - [func LookupHost(host string) (addrs []string, err error)](#func-lookuphost)
    - [func LookupPort(network, service string) (port int, err error)](#func-lookupport)
    - [func LookupTXT(name string) ([]string, error)](#func-lookuptxt)
    - [func ParseCIDR(s string) (IP, *IPNet, error)](#func-parsecidr)
    - [func Pipe() (Conn, Conn)](#func-pipe)
    - [func SplitHostPort(hostport string) (host, port string, err error)](#func-splithostport)
- [Types](#类型) 
    - [type Addr]()
        - [func InterfaceAddrs()([]Addr, error)]()
    - [type AddrError]()
        - [func (e *AddrError) Error() string]()
        - [func (e *AddrError) Temporary() bool]()
        - [func (e *AddrError) Timeout() bool]()
    - [type Buffers]()
        - []()
        - []()
    - [type Conn]()
        - []()
        - []()
        - []()
    - [type DNSConfigError]()
        - []()
        - []()
        - []()
        - []()
    - [type DNSError]()
        - []()
        - []()
        - []()
    - [type Dialer]()
        - []()
        - []()
    - [type Error]()
    - [type Flags]()
        - []()
    - [type HardwareAddr]()
        - []()
        - []()
    - [type IP]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type IPAddr]()
        - []()
        - []()
        - []()
    - [type IPConn]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type IPMask]()
        - []()
        - []()
        - []()
        - []()
    - [type IPNet]()
        - []()
        - []()
        - []()
    - [type Interface]()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type InvalidAddrError]()
        - []()
        - []()
        - []()
    - [type ListenConfig]()
        - []()
        - []()
    - [type Listener]()
        - []()
        - []()
    - [type MX]()
        - []()
    - [type NS]()
        - []()
    - [type OpError]()
        - []()
        - []()
        - []()
        - []()
    - [type PacketConn]()
        - []()
        - []()
    - [type ParseError]()
        - []()
        - []()
        - []()
    - [type Resolver]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type SRV]()
        - []()
    - [type TCPAddr]()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type TCPConn]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type TCPListener]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type UDPAddr]()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type UDPConn]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type UnixAddr]()
        - []()
        - []()
        - []()
    - [type UnixConn]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type UnixListener]()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
        - []()
    - [type UnknownNetworkError]()
        - []()
        - []()
        - []()
    - [Bugs]()

## 常量

IP 地址长度（字节）。

```go
const (
    IPv4len = 4
    IPv6len = 16
)
```

## 变量

熟知的 IPv4 地址

```go
var (
	IPv4bcast     = IPv4(255, 255, 255, 255) // limited broadcast
	IPv4allsys    = IPv4(224, 0, 0, 1)       // all systems
	IPv4allrouter = IPv4(224, 0, 0, 2)       // all routers
	IPv4zero      = IPv4(0, 0, 0, 0)         // all zeros
)
```

熟知的 IPv6 地址

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

## 函数

### func JoinHostPort

```go
func JoinHostPort(host, port string) string
```

JoinHostPort 将主机和端口组合成 “host:port” 形式的网络地址。如果 host 包含冒号，如在文字 IPv6 地址中找到的那样，则 JoinHostPort 返回 “[host]:port”。

有关主机和端口参数的说明，请参见 [func Dial](#func-dial)。

### func LookupAddr

```go
func LookupAddr(addr string) (names []string, err error)
```

LookupAddr 对给定地址执行反向查找，返回映射到该地址的名称列表。

返回的名称被验证为格式正确的表示格式域名。如果响应包含无效名称，则这些记录将被过滤掉，并且将与剩余结果（如果有）一起返回错误。

使用宿主 C 库解析器时，最多返回一个结果。要绕过主机解析器，请使用自定义解析器。

LookupAddr 内部使用 context.Background；要指定上下文，请使用 Resolver.LookupAddr。

### func LookupCNAME

```go
func LookupCNAME(host string) (cname string, err error)
```

LookupCNAME 返回给定主机的规范名称。不关心规范名称的调用者可以直接调用 LookupHost 或 LookupIP；两者都负责将规范名称解析为查找的一部分。

规范名称是跟随零个或多个 CNAME 记录后的最终名称。如果主机不包含 DNS“CNAME”记录，只要主机解析为地址记录，LookupCNAME 就不会返回错误。

返回的规范名称被验证为格式正确的表示格式域名。

LookupCNAME 在内部使用 context.Background；要指定上下文，请使用 Resolver.LookupCNAME。

### func LookupHost

```go
func LookupHost(host string) (addrs []string, err error)
```

LookupHost 使用本地解析器查找给定的主机。它返回该主机地址的一部分。

LookupHost 内部使用 context.Background；要指定上下文，请使用 Resolver.LookupHost。

### func LookupPort

```go
func LookupPort(network, service string) (port int, err error)
```

LookupPort 查找给定网络和服务的端口。

LookupPort 内部使用 context.Background；要指定上下文，请使用 Resolver.LookupPort。

### func LookupTXT

```go
func LookupTXT(name string) ([]string, error)
```

LookupTXT 返回给定域名的 DNS TXT 记录。

LookupTXT 内部使用 context.Background；要指定上下文，请使用 Resolver.LookupTXT。

### func ParseCIDR

```go
func ParseCIDR(s string) (IP, *IPNet, error)
```

ParseCIDR 将 s 解析为 CIDR 表示法 IP 地址和前缀长度，如“192.0.2.0/24”或“2001:db8::/32”，如 RFC 4632 和 RFC 4291 中所定义。

它返回 IP 地址以及 IP 和前缀长度隐含的网络。例如，ParseCIDR("192.0.2.1/24") 返回 IP 地址 192.0.2.1 和网络 192.0.2.0/24。

Example

```go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)

}
```

### func Pipe

```go
func Pipe() (Conn, Conn)
```

Pipe 创建一个同步的、内存中的、全双工的网络连接；两端都实现了Conn接口。一端的读取与另一端的写入匹配，直接在两者之间复制数据；没有内部缓冲。

### func SplitHostPort

```go
func SplitHostPort(hostport string) (host, port string, err error)
```

SplitHostPort 将“host:port”、“host%zone:port”、“[host]:port”或“[host%zone]:port”形式的网络地址拆分为 host 或 host%zone 和端口。

hostport 中的文字 IPv6 地址必须括在方括号中，如 “[::1]:80”、“[::1%lo0]:80”。

有关 hostport 参数以及主机和端口结果的说明，请参见 [func Dial](#func-dial)。

## 类型

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


## Bugs

- 在 JS 和 Windows 上，未实现 FileConn、FileListener 和 FilePacketConn 函数。
- 在JS上，并没有实现与 Interface 相关的方法和功能。
- 在 AIX、DragonFly BSD、NetBSD、OpenBSD、Plan 9 和 Solaris 上，未实现接口的 MulticastAddrs 方法。
- 在每个 POSIX 平台上，使用 ReadFrom 或 ReadFromIP 方法从“ip4”网络读取可能不会返回完整的 IPv4 数据包，包括其标头，即使有可用空间也是如此。即使在 Read 或 ReadMsgIP 可以返回完整数据包的情况下，也会发生这种情况。因此，如果接收完整数据包很重要，建议您不要使用这些方法。
- Go 1 兼容性指南使我们无法更改这些方法的行为；请改用 Read 或 ReadMsgIP。
- 在 JS 和 Plan 9 上，没有实现 IPConn 相关的方法和函数。
- 在 Windows 上，未实现 IPConn 的 File 方法。
- 在 DragonFly BSD 和 OpenBSD 上，侦 听“tcp” 和 “udp” 网络不会同时侦听 IPv4 和 IPv6 连接。这是因为 IPv4 流量不会路由到 IPv6 套接字 - 如果要支持两个地址系列，则需要两个单独的套接字。有关详细信息，请参阅 inet6(4)。
- 在 Windows 上，syscall.RawConn 的 Write 方法不与运行时的网络轮询器集成。它不能等待连接变得可写，并且不遵守截止日期。如果用户提供的回调返回 false，Write 方法将立即失败。
- 在 JS 和 Plan 9 上，syscall.RawConn 的 Control、Read 和 Write 方法没有实现。
- 在 JS 和 Windows 上，未实现 TCPConn 和 TCPListener 的 File 方法。
- 在 Plan 9 上，UDPConn 的 ReadMsgUDP 和 WriteMsgUDP 方法没有实现。
- 在 Windows 上，未实现 UDPConn 的 File 方法。
- 在 JS 上，没有实现 UDPConn 相关的方法和函数。
- 在 JS 和 Plan 9 上，没有实现 UnixConn 和 UnixListener 相关的方法和函数。
- 在 Windows 上，与 UnixConn 和 UnixListener 相关的方法和函数不适用于 “unixgram” 和 “unixpacket”。

## 目录

- [http](net/http.md) 
    - [cgi](net/http/cgi.md)
    - [cookiejar](net/http/cookiejar.md)
    - [fcgi](net/http/fcgi.md)
    - [httptest](net/http/httptest.md)
    - [httptrace](net/http/httptrace.md)
    - [httputil](net/http/httputil.md)
    - [pprof](net/http/pprof.md)
- [mail](net/mail.md)
- [netip](net/netip.md)
- [rpc](net/rpc.md)
    - [jsonrpc](net/rpc/jsonrpc.md)
- [smtp](net/smtp.md)
- [textproto](net/textproto.md)
- [url](net/url.md)
- internal
    - [socktest](net/internal/socktest.md)