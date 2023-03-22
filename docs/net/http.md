
# http

## 概述

http 包提供 HTTP 客户端和服务器实现。

Get、Head、Post 和 PostForm 发出 HTTP（或 HTTPS）请求：

```go
resp, err := http.Get("http://example.com")
···
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
···
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
```

客户端必须在完成后关闭响应主体：

```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
// ..
```

要控制 HTTP 客户端标头、重定向策略和其他设置，请创建一个 Client：

```go
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

要控制代理、TLS 配置、保持活动、压缩和其他设置，请创建一个 Transport：

```go
tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

Clients 和 Transports 对于多个 goroutines 的并发使用是安全的，并且为了效率应该只创建一次并重复使用。

ListenAndServe 使用给定的地址和处理程序启动 HTTP 服务器。 handler 通常为 nil，表示使用 DefaultServeMux。 Handle 和 HandleFunc 将处理程序添加到 DefaultServeMux：

```go
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

通过创建自定义 Server 可以更好地控制服务器的行为：

```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

从 Go 1.6 开始，http 包在使用 HTTPS 时透明支持 HTTP/2 协议。必须禁用 HTTP/2 的程序可以通过将 Transport.TLSNextProto（对于客户端）或 Server.TLSNextProto（对于服务器）设置为非零的空映射来实现。或者，目前支持以下 GODEBUG 环境变量：

```bash
GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps
```

GODEBUG 变量不在 Go 的 API 兼容性承诺范围内。请在禁用 HTTP/2 支持之前报告任何问题：https://golang.org/s/http2bug

http 包的 Transport 和 Server 都自动启用 HTTP/2 支持以进行简单配置。要为更复杂的配置启用 HTTP/2，使用较低级别的 HTTP/2 功能，或使用更新版本的 Go 的 http2 包，直接导入“golang.org/x/net/http2”并使用其 ConfigureTransport 和/或 ConfigureServer 功能。通过 golang.org/x/net/http2 包手动配置 HTTP/2 优先于 net/http 包的内置 HTTP/2 支持。

## 索引

- [Constants](#常量)
- [Variables](#变量)
- [Functions](#函数)
- [Types](#类型)

## 常量

```go
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)
```

常见的 HTTP 方法。

除非另有说明，否则这些在 RFC 7231 第 4.3 节中定义。

```go
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	StatusFound            = 302 // RFC 9110, 15.4.3
	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)
```

在 IANA 注册的 HTTP 状态代码。请参阅：https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml


```go
const DefaultMaxHeaderBytes = 1 << 20 // 1 MB
```

DefaultMaxHeaderBytes 是 HTTP 请求中标头的最大允许大小。这可以通过设置 Server.MaxHeaderBytes 来覆盖。

```go
const DefaultMaxIdleConnsPerHost = 2
```

DefaultMaxIdleConnsPerHost 是 Transport 的 MaxIdleConnsPerHost 的默认值。

```go
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
```

TimeFormat 是在 HTTP 标头中生成时间时使用的时间格式。它类似于 time.RFC1123，但将 GMT 硬编码为时区。格式化的时间必须采用 UTC 格式才能生成正确的格式。

解析这种时间格式，参见 [ParseTime](#func-parsetime)。

```go
const TrailerPrefix = "Trailer:"
```

TrailerPrefix 是 ResponseWriter.Header 映射键的魔法前缀，如果存在，表示映射条目实际上是用于响应尾部，而不是响应头。在 ServeHTTP 调用完成并且值在预告片中发送后，前缀将被剥离。

此机制仅适用于在写入标头之前未知的尾部。如果尾部集合在写入头之前是固定的或已知的，则首选普通的 Go 尾部机制：

```html
https://pkg.go.dev/net/http#ResponseWriter
https://pkg.go.dev/net/http#example-ResponseWriter-Trailers
```

## 变量

```go
var (
	// ErrNotSupported indicates that a feature is not supported.
	//
	// It is returned by ResponseController methods to indicate that
	// the handler does not support the method, and by the Push method
	// of Pusher implementations to indicate that HTTP/2 Push support
	// is not available.
	ErrNotSupported = &ProtocolError{"feature not supported"}

	// Deprecated: ErrUnexpectedTrailer is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrUnexpectedTrailer = &ProtocolError{"trailer header without chunked transfer encoding"}

	// ErrMissingBoundary is returned by Request.MultipartReader when the
	// request's Content-Type does not include a "boundary" parameter.
	ErrMissingBoundary = &ProtocolError{"no multipart boundary param in Content-Type"}

	// ErrNotMultipart is returned by Request.MultipartReader when the
	// request's Content-Type is not multipart/form-data.
	ErrNotMultipart = &ProtocolError{"request Content-Type isn't multipart/form-data"}

	// Deprecated: ErrHeaderTooLong is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrHeaderTooLong = &ProtocolError{"header too long"}

	// Deprecated: ErrShortBody is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrShortBody = &ProtocolError{"entity body too short"}

	// Deprecated: ErrMissingContentLength is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrMissingContentLength = &ProtocolError{"missing ContentLength in HEAD response"}
)
```

```go
var (
	// ErrBodyNotAllowed is returned by ResponseWriter.Write calls
	// when the HTTP method or response code does not permit a
	// body.
	ErrBodyNotAllowed = errors.New("http: request method or response status code does not allow body")

	// ErrHijacked is returned by ResponseWriter.Write calls when
	// the underlying connection has been hijacked using the
	// Hijacker interface. A zero-byte write on a hijacked
	// connection will return ErrHijacked without any other side
	// effects.
	ErrHijacked = errors.New("http: connection has been hijacked")

	// ErrContentLength is returned by ResponseWriter.Write calls
	// when a Handler set a Content-Length response header with a
	// declared size and then attempted to write more bytes than
	// declared.
	ErrContentLength = errors.New("http: wrote more than the declared Content-Length")

	// Deprecated: ErrWriteAfterFlush is no longer returned by
	// anything in the net/http package. Callers should not
	// compare errors against this variable.
	ErrWriteAfterFlush = errors.New("unused")
)
```

HTTP 服务器使用的错误。

```go
var (
	// ServerContextKey is a context key. It can be used in HTTP
	// handlers with Context.Value to access the server that
	// started the handler. The associated value will be of
	// type *Server.
	ServerContextKey = &contextKey{"http-server"}

	// LocalAddrContextKey is a context key. It can be used in
	// HTTP handlers with Context.Value to access the local
	// address the connection arrived on.
	// The associated value will be of type net.Addr.
	LocalAddrContextKey = &contextKey{"local-addr"}
)
```

```go
var DefaultClient = &Client{}
```

DefaultClient 是默认 Client，由 Get、Head 和 Post 使用。

```go
var DefaultServeMux = &defaultServeMux
```

DefaultServeMux 是 Serve 使用的默认 ServeMux。

```go
var ErrAbortHandler = errors.New("net/http: abort Handler")
```

ErrAbortHandler 是一个用于中止处理程序的哨兵恐慌值。虽然来自 ServeHTTP 的任何恐慌都会中止对客户端的响应，但使用 ErrAbortHandler 恐慌也会抑制将堆栈跟踪记录到服务器的错误日志中。

```go
var ErrBodyReadAfterClose = errors.New("http: invalid Read on closed Body")
```

在主体关闭后读取请求或响应主体时返回 ErrBodyReadAfterClose。这通常发生在 HTTP 处理程序在其 ResponseWriter 上调用 WriteHeader 或 Write 后读取正文时。

```go
var ErrHandlerTimeout = errors.New("http: Handler timeout")
```

ErrHandlerTimeout 在已超时的处理程序中的 ResponseWriter Write 调用上返回。

```go
var ErrLineTooLong = internal.ErrLineTooLong
```

读取具有错误分块编码的请求或响应主体时返回 ErrLineTooLong。

```go
var ErrMissingFile = errors.New("http: no such file")
```

当提供的文件字段名称在请求中不存在或不是文件字段时，FormFile 将返回 ErrMissingFile。

```go
var ErrNoCookie = errors.New("http: named cookie not present")
```

当找不到 cookie 时，Request 的 Cookie 方法会返回 ErrNoCookie。

```go
var ErrNoLocation = errors.New("http: no Location header in response")
```

当没有 Location 标头存在时，Response 的 Location 方法返回 ErrNoLocation。

```go
var ErrServerClosed = errors.New("http: Server closed")
```

在调用 Shutdown 或 Close 后，服务器的 Serve、ServeTLS、ListenAndServe 和 ListenAndServeTLS 方法返回 ErrServerClosed。

```go
var ErrSkipAltProtocol = errors.New("net/http: skip alternate protocol")
```

ErrSkipAltProtocol 是由 Transport.RegisterProtocol 定义的哨兵错误值。

```go
var ErrUseLastResponse = errors.New("net/http: use last response")
```

Client.CheckRedirect 挂钩可以返回 ErrUseLastResponse 以控制如何处理重定向。如果返回，则不发送下一个请求，并返回最新的响应，其正文未关闭。

```go
var NoBody = noBody{}
```

NoBody 是一个没有字节的 io.ReadCloser。 Read 总是返回 EOF 而 Close 总是返回 nil。它可以在传出的客户端请求中使用，以明确表示请求的字节数为零。然而，另一种方法是简单地将 Request.Body 设置为 nil。

## 函数

### func CanonicalHeaderKey

```go

```

### func DetectContentType

```go

```

### func Error

```go

```

### func Handle

```go

```

### func HandleFunc

```go

```

### func ListenAndServe

```go

```

### func ListenAndServeTLS

```go

```

### func MaxBytesReader

```go

```

### func NotFound

```go

```

### func ParseHTTPVersion

```go

```

### func ParseTime

```go

```

### func ProxyFromEnvironment

```go

```

### func ProxyURL

```go

```

### func Redirect

```go

```

### func Serve

```go

```

### func ServeContent

```go

```

### func ServeFile

```go

```

### func ServeTLS

```go

```

### func SetCookie

```go

```

### func StatusText

```go

```

## 类型

### type Client

```go
type Client struct {
	// Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	Transport RoundTripper

	// CheckRedirect specifies the policy for handling redirects.
	// If CheckRedirect is not nil, the client calls it before
	// following an HTTP redirect. The arguments req and via are
	// the upcoming request and the requests made already, oldest
	// first. If CheckRedirect returns an error, the Client's Get
	// method returns both the previous Response (with its Body
	// closed) and CheckRedirect's error (wrapped in a url.Error)
	// instead of issuing the Request req.
	// As a special case, if CheckRedirect returns ErrUseLastResponse,
	// then the most recent response is returned with its body
	// unclosed, along with a nil error.
	//
	// If CheckRedirect is nil, the Client uses its default policy,
	// which is to stop after 10 consecutive requests.
	CheckRedirect func(req *Request, via []*Request) error

	// Jar specifies the cookie jar.
	//
	// The Jar is used to insert relevant cookies into every
	// outbound Request and is updated with the cookie values
	// of every inbound Response. The Jar is consulted for every
	// redirect that the Client follows.
	//
	// If Jar is nil, cookies are only sent if they are explicitly
	// set on the Request.
	Jar CookieJar

	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	//
	// The Client cancels requests to the underlying Transport
	// as if the Request's Context ended.
	//
	// For compatibility, the Client will also use the deprecated
	// CancelRequest method on Transport if found. New
	// RoundTripper implementations should use the Request's Context
	// for cancellation instead of implementing CancelRequest.
	Timeout time.Duration
}
```

### func (*Client) CloseIdleConnections

```go

```

### func (*Client) Do

```go

```

### func (*Client) Get

```go

```

### func (*Client) Head

```go

```

### func (*Client) Post

```go

```

### func (*Client) PostForm

```go

```

### type ConnState

```go
type ConnState int
```

### func (ConnState) String

```go

```

### type Cookie

```go
type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
```

Cookie 表示在 HTTP 响应的 Set-Cookie 标头或 HTTP 请求的 Cookie 标头中发送的 HTTP cookie。

有关详细信息，请参阅 https://tools.ietf.org/html/rfc6265。

### func (*Cookie) String

```go

```

### func (*Cookie) Valid

```go

```

### type CookieJar

```go

```

### type Dir

```go
type Dir string
```

### func (Dir) Open

```go

```

### type File

```go
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}
```

### type FileSystem

```go
type FileSystem interface {
	Open(name string) (File, error)
}
```

### func FS

```go

```

### type Flusher

```go
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}
```

### type Handler

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

### func AllowQuerySemicolons

```go
func AllowQuerySemicolons(h Handler) Handler
```

### func FileServer

```go

```

### func MaxBytesHandler

```go

```

### func NotFoundHandler

```go

```

### func RedirectHandler

```go

```

### func StripPrefix

```go

```

### func TimeoutHandler

```go

```

### type HandlerFunc

```go
type HandlerFunc func(ResponseWriter, *Request)
```

HandlerFunc 类型是一个适配器，允许将普通函数用作 HTTP 处理程序。如果 f 是具有适当签名的函数，则 HandlerFunc(f) 是调用 f 的处理程序。

### func (HandlerFunc) ServeHTTP

```go

```

### type Header

```go
type Header map[string][]string
```

Header 表示 HTTP 标头中的键值对。

密钥应采用规范形式，如 CanonicalHeaderKey 返回的那样。

### func (Header) Add

```go

```

### func (Header) Clone

```go

```

### func (Header) Get

```go

```

### func (Header) Set

```go

```

### func (Header) Values

```go

```

### func (Header) Write

```go

```

### func (Header) WriteSubset

```go

```

### type Hijacker

```go
type Hijacker interface {
	// Hijack lets the caller take over the connection.
	// After a call to Hijack the HTTP server library
	// will not do anything else with the connection.
	//
	// It becomes the caller's responsibility to manage
	// and close the connection.
	//
	// The returned net.Conn may have read or write deadlines
	// already set, depending on the configuration of the
	// Server. It is the caller's responsibility to set
	// or clear those deadlines as needed.
	//
	// The returned bufio.Reader may contain unprocessed buffered
	// data from the client.
	//
	// After a call to Hijack, the original Request.Body must not
	// be used. The original Request's Context remains valid and
	// is not canceled until the Request's ServeHTTP method
	// returns.
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```

### type MaxBytesError

```go
type MaxBytesError struct {
	Limit int64
}
```

当超过读取限制时，MaxBytesReader 会返回 MaxBytesError。

### func (*MaxBytesError) Error

```go

```

### type ProtocolError (DEPRECATED)

```go
type ProtocolError struct {
	ErrorString string
}
```

### type PushOptions

```go
type PushOptions struct {
	// Method specifies the HTTP method for the promised request.
	// If set, it must be "GET" or "HEAD". Empty means "GET".
	Method string

	// Header specifies additional promised request headers. This cannot
	// include HTTP/2 pseudo header fields like ":path" and ":scheme",
	// which will be added automatically.
	Header Header
}
```

PushOptions 描述了 Pusher.Push 的选项。

### type Pusher

```go
type Pusher interface {
	// Push initiates an HTTP/2 server push. This constructs a synthetic
	// request using the given target and options, serializes that request
	// into a PUSH_PROMISE frame, then dispatches that request using the
	// server's request handler. If opts is nil, default options are used.
	//
	// The target must either be an absolute path (like "/path") or an absolute
	// URL that contains a valid host and the same scheme as the parent request.
	// If the target is a path, it will inherit the scheme and host of the
	// parent request.
	//
	// The HTTP/2 spec disallows recursive pushes and cross-authority pushes.
	// Push may or may not detect these invalid pushes; however, invalid
	// pushes will be detected and canceled by conforming clients.
	//
	// Handlers that wish to push URL X should call Push before sending any
	// data that may trigger a request for URL X. This avoids a race where the
	// client issues requests for X before receiving the PUSH_PROMISE for X.
	//
	// Push will run in a separate goroutine making the order of arrival
	// non-deterministic. Any required synchronization needs to be implemented
	// by the caller.
	//
	// Push returns ErrNotSupported if the client has disabled push or if push
	// is not supported on the underlying connection.
	Push(target string, opts *PushOptions) error
}
```

Pusher 是 ResponseWriters 实现的接口，支持 HTTP/2 服务器推送。有关更多背景信息，请参阅 https://tools.ietf.org/html/rfc7540#section-8.2。

### type Request

```go
type Request struct {
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	// For client requests, an empty string means GET.
	//
	// Go's HTTP client does not support sending a request with
	// the CONNECT method. See the documentation on Transport for
	// details.
	Method string

	// URL specifies either the URI being requested (for server
	// requests) or the URL to access (for client requests).
	//
	// For server requests, the URL is parsed from the URI
	// supplied on the Request-Line as stored in RequestURI.  For
	// most requests, fields other than Path and RawQuery will be
	// empty. (See RFC 7230, Section 5.3)
	//
	// For client requests, the URL's Host specifies the server to
	// connect to, while the Request's Host field optionally
	// specifies the Host header value to send in the HTTP
	// request.
	URL *url.URL

	// The protocol version for incoming server requests.
	//
	// For client requests, these fields are ignored. The HTTP
	// client code always uses either HTTP/1.1 or HTTP/2.
	// See the docs on Transport for details.
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0

	// Header contains the request header fields either received
	// by the server or to be sent by the client.
	//
	// If a server received a request with header lines,
	//
	//	Host: example.com
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	fOO: Bar
	//	foo: two
	//
	// then
	//
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
	//
	// For incoming requests, the Host header is promoted to the
	// Request.Host field and removed from the Header map.
	//
	// HTTP defines that header names are case-insensitive. The
	// request parser implements this by using CanonicalHeaderKey,
	// making the first character and any characters following a
	// hyphen uppercase and the rest lowercase.
	//
	// For client requests, certain headers such as Content-Length
	// and Connection are automatically written when needed and
	// values in Header may be ignored. See the documentation
	// for the Request.Write method.
	Header Header

	// Body is the request's body.
	//
	// For client requests, a nil body means the request has no
	// body, such as a GET request. The HTTP Client's Transport
	// is responsible for calling the Close method.
	//
	// For server requests, the Request Body is always non-nil
	// but will return EOF immediately when no body is present.
	// The Server will close the request body. The ServeHTTP
	// Handler does not need to.
	//
	// Body must allow Read to be called concurrently with Close.
	// In particular, calling Close should unblock a Read waiting
	// for input.
	Body io.ReadCloser

	// GetBody defines an optional func to return a new copy of
	// Body. It is used for client requests when a redirect requires
	// reading the body more than once. Use of GetBody still
	// requires setting Body.
	//
	// For server requests, it is unused.
	GetBody func() (io.ReadCloser, error)

	// ContentLength records the length of the associated content.
	// The value -1 indicates that the length is unknown.
	// Values >= 0 indicate that the given number of bytes may
	// be read from Body.
	//
	// For client requests, a value of 0 with a non-nil Body is
	// also treated as unknown.
	ContentLength int64

	// TransferEncoding lists the transfer encodings from outermost to
	// innermost. An empty list denotes the "identity" encoding.
	// TransferEncoding can usually be ignored; chunked encoding is
	// automatically added and removed as necessary when sending and
	// receiving requests.
	TransferEncoding []string

	// Close indicates whether to close the connection after
	// replying to this request (for servers) or after sending this
	// request and reading its response (for clients).
	//
	// For server requests, the HTTP server handles this automatically
	// and this field is not needed by Handlers.
	//
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	Close bool

	// For server requests, Host specifies the host on which the
	// URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
	// is either the value of the "Host" header or the host name
	// given in the URL itself. For HTTP/2, it is the value of the
	// ":authority" pseudo-header field.
	// It may be of the form "host:port". For international domain
	// names, Host may be in Punycode or Unicode form. Use
	// golang.org/x/net/idna to convert it to either format if
	// needed.
	// To prevent DNS rebinding attacks, server Handlers should
	// validate that the Host header has a value for which the
	// Handler considers itself authoritative. The included
	// ServeMux supports patterns registered to particular host
	// names and thus protects its registered Handlers.
	//
	// For client requests, Host optionally overrides the Host
	// header to send. If empty, the Request.Write method uses
	// the value of URL.Host. Host may contain an international
	// domain name.
	Host string

	// Form contains the parsed form data, including both the URL
	// field's query parameters and the PATCH, POST, or PUT form data.
	// This field is only available after ParseForm is called.
	// The HTTP client ignores Form and uses Body instead.
	Form url.Values

	// PostForm contains the parsed form data from PATCH, POST
	// or PUT body parameters.
	//
	// This field is only available after ParseForm is called.
	// The HTTP client ignores PostForm and uses Body instead.
	PostForm url.Values

	// MultipartForm is the parsed multipart form, including file uploads.
	// This field is only available after ParseMultipartForm is called.
	// The HTTP client ignores MultipartForm and uses Body instead.
	MultipartForm *multipart.Form

	// Trailer specifies additional headers that are sent after the request
	// body.
	//
	// For server requests, the Trailer map initially contains only the
	// trailer keys, with nil values. (The client declares which trailers it
	// will later send.)  While the handler is reading from Body, it must
	// not reference Trailer. After reading from Body returns EOF, Trailer
	// can be read again and will contain non-nil values, if they were sent
	// by the client.
	//
	// For client requests, Trailer must be initialized to a map containing
	// the trailer keys to later send. The values may be nil or their final
	// values. The ContentLength must be 0 or -1, to send a chunked request.
	// After the HTTP request is sent the map values can be updated while
	// the request body is read. Once the body returns EOF, the caller must
	// not mutate Trailer.
	//
	// Few HTTP clients, servers, or proxies support HTTP trailers.
	Trailer Header

	// RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. This field is not filled in by ReadRequest and
	// has no defined format. The HTTP server in this package
	// sets RemoteAddr to an "IP:port" address before invoking a
	// handler.
	// This field is ignored by the HTTP client.
	RemoteAddr string

	// RequestURI is the unmodified request-target of the
	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
	// to a server. Usually the URL field should be used instead.
	// It is an error to set this field in an HTTP client request.
	RequestURI string

	// TLS allows HTTP servers and other software to record
	// information about the TLS connection on which the request
	// was received. This field is not filled in by ReadRequest.
	// The HTTP server in this package sets the field for
	// TLS-enabled connections before invoking a handler;
	// otherwise it leaves the field nil.
	// This field is ignored by the HTTP client.
	TLS *tls.ConnectionState

	// Cancel is an optional channel whose closure indicates that the client
	// request should be regarded as canceled. Not all implementations of
	// RoundTripper may support Cancel.
	//
	// For server requests, this field is not applicable.
	//
	// Deprecated: Set the Request's context with NewRequestWithContext
	// instead. If a Request's Cancel field and context are both
	// set, it is undefined whether Cancel is respected.
	Cancel <-chan struct{}

	// Response is the redirect response which caused this request
	// to be created. This field is only populated during client
	// redirects.
	Response *Response
	// contains filtered or unexported fields
}
```

Request 表示由服务器接收或由客户端发送的 HTTP 请求。

客户端和服务器用法之间的字段语义略有不同。除了以下字段的注释外，请参阅 Request.Write 和 RoundTripper 的文档。

### func NewRequest

```go

```

### func NewRequestWithContext

```go

```

### func ReadRequest

```go

```

### func (*Request) AddCookie

```go

```

### func (*Request) BasicAuth

```go

```

### func (*Request) Clone

```go

```

### func (*Request) Context

```go

```

### func (*Request) Cookie

```go

```

### func (*Request) Cookies

```go

```

### func (*Request) FormFile

```go

```

### func (*Request) FormValue

```go

```

### func (*Request) MultipartReader

```go

```

### func (*Request) ParseForm

```go

```

### func (*Request) ParseMultipartForm

```go

```

### func (*Request) PostFormValue

```go

```

### func (*Request) ProtoAtLeast

```go

```

### func (*Request) Referer

```go

```

### func (*Request) SetBasicAuth

```go

```

### func (*Request) UserAgent

```go

```

### func (*Request) WithContext

```go

```

### func (*Request) Write

```go

```

### func (*Request) WriteProxy

```go

```

### type Response

```go
type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header Header

	// Body represents the response body.
	//
	// The response body is streamed on demand as the Body field
	// is read. If the network connection fails or the server
	// terminates the response, Body.Read calls return an error.
	//
	// The http Client and Transport guarantee that Body is always
	// non-nil, even on responses without a body or responses with
	// a zero-length body. It is the caller's responsibility to
	// close Body. The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	//
	// The Body is automatically dechunked if the server replied
	// with a "chunked" Transfer-Encoding.
	//
	// As of Go 1.12, the Body will also implement io.Writer
	// on a successful "101 Switching Protocols" response,
	// as used by WebSockets and HTTP/2's "h2c" mode.
	Body io.ReadCloser

	// ContentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	ContentLength int64

	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	TransferEncoding []string

	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	Close bool

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	Uncompressed bool

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	Trailer Header

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	Request *Request

	// TLS contains information about the TLS connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	TLS *tls.ConnectionState
}
```

Response 表示来自 HTTP 请求的响应。

一旦收到响应头，客户端和传输就从服务器返回响应。在读取 Body 字段时，响应主体按需流式传输。

### func Get

```go

```

### func Head

```go

```

### func Post

```go

```

### func PostForm

```go

```

### func ReadResponse

```go

```

### func (*Response) Cookies

```go

```

### func (*Response) Location

```go

```

### func (*Response) ProtoAtLeast

```go

```

### func (*Response) Write

```go

```

### type ResponseController

```go
type ResponseController struct {
	// contains filtered or unexported fields
}
```

HTTP 处理程序使用 ResponseController 来控制响应。

在 Handler.ServeHTTP 方法返回后，不得使用 ResponseController。

### func NewResponseController

```go

```

### func (*ResponseController) Flush

```go

```

### func (*ResponseController) Hijack
```go

```

### func (*ResponseController) SetReadDeadline

```go

```

### func (*ResponseController) SetWriteDeadline

```go

```

### type ResponseWriter

```go
type ResponseWriter interface {
	// Header returns the header map that will be sent by
	// WriteHeader. The Header map also is the mechanism with which
	// Handlers can set HTTP trailers.
	//
	// Changing the header map after a call to WriteHeader (or
	// Write) has no effect unless the HTTP status code was of the
	// 1xx class or the modified headers are trailers.
	//
	// There are two ways to set Trailers. The preferred way is to
	// predeclare in the headers which trailers you will later
	// send by setting the "Trailer" header to the names of the
	// trailer keys which will come later. In this case, those
	// keys of the Header map are treated as if they were
	// trailers. See the example. The second way, for trailer
	// keys not known to the Handler until after the first Write,
	// is to prefix the Header map keys with the TrailerPrefix
	// constant value. See TrailerPrefix.
	//
	// To suppress automatic response headers (such as "Date"), set
	// their value to nil.
	Header() Header

	// Write writes the data to the connection as part of an HTTP reply.
	//
	// If WriteHeader has not yet been called, Write calls
	// WriteHeader(http.StatusOK) before writing the data. If the Header
	// does not contain a Content-Type line, Write adds a Content-Type set
	// to the result of passing the initial 512 bytes of written data to
	// DetectContentType. Additionally, if the total size of all written
	// data is under a few KB and there are no Flush calls, the
	// Content-Length header is added automatically.
	//
	// Depending on the HTTP protocol version and the client, calling
	// Write or WriteHeader may prevent future reads on the
	// Request.Body. For HTTP/1.x requests, handlers should read any
	// needed request body data before writing the response. Once the
	// headers have been flushed (due to either an explicit Flusher.Flush
	// call or writing enough data to trigger a flush), the request body
	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
	// handlers to continue to read the request body while concurrently
	// writing the response. However, such behavior may not be supported
	// by all HTTP/2 clients. Handlers should read before writing if
	// possible to maximize compatibility.
	Write([]byte) (int, error)

	// WriteHeader sends an HTTP response header with the provided
	// status code.
	//
	// If WriteHeader is not called explicitly, the first call to Write
	// will trigger an implicit WriteHeader(http.StatusOK).
	// Thus explicit calls to WriteHeader are mainly used to
	// send error codes or 1xx informational responses.
	//
	// The provided code must be a valid HTTP 1xx-5xx status code.
	// Any number of 1xx headers may be written, followed by at most
	// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
	// headers may be buffered. Use the Flusher interface to send
	// buffered data. The header map is cleared when 2xx-5xx headers are
	// sent, but not with 1xx headers.
	//
	// The server will automatically send a 100 (Continue) header
	// on the first read from the request body if the request has
	// an "Expect: 100-continue" header.
	WriteHeader(statusCode int)
}
```

HTTP 处理程序使用 ResponseWriter 接口来构造 HTTP 响应。

在 Handler.ServeHTTP 方法返回后，不得使用 ResponseWriter。

### type RoundTripper

```go
type RoundTripper interface {
	// RoundTrip executes a single HTTP transaction, returning
	// a Response for the provided Request.
	//
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// A non-nil err should be reserved for failure to obtain a
	// response. Similarly, RoundTrip should not attempt to
	// handle higher-level protocol details such as redirects,
	// authentication, or cookies.
	//
	// RoundTrip should not modify the request, except for
	// consuming and closing the Request's Body. RoundTrip may
	// read fields of the request in a separate goroutine. Callers
	// should not mutate or reuse the request until the Response's
	// Body has been closed.
	//
	// RoundTrip must always close the body, including on errors,
	// but depending on the implementation may do so in a separate
	// goroutine even after RoundTrip returns. This means that
	// callers wanting to reuse the body for subsequent requests
	// must arrange to wait for the Close call before doing so.
	//
	// The Request's URL and Header fields must be initialized.
	RoundTrip(*Request) (*Response, error)
}
```

### func NewFileTransport

```go

```

### type SameSite

```go
type SameSite int
```

SameSite 允许服务器定义 cookie 属性，使浏览器无法将此 cookie 与跨站点请求一起发送。主要目标是减轻跨源信息泄露的风险，并提供一些保护措施来防止跨站点请求伪造攻击。

有关详细信息，请参阅 https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00。

### type ServeMux

```go
type ServeMux struct {
	// contains filtered or unexported fields
}
```

ServeMux 是一个 HTTP 请求多路复用器。它将每个传入请求的 URL 与已注册模式列表进行匹配，并为与 URL 最匹配的模式调用处理程序。

模式命名固定的根路径，如“/favicon.ico”，或根子树，如“/images/”（注意尾部斜杠）。较长的模式优先于较短的模式，因此如果有为 “/images/” 和 “/images/thumbnails/” 注册的处理程序，则后者的处理程序将被调用以开始于 “/images/thumbnails/” 的路径，而前者将接收对 “/images/” 子树中任何其他路径的请求。

请注意，由于以斜杠结尾的模式命名为有根子树，因此模式 “/” 匹配所有其他已注册模式不匹配的路径，而不仅仅是 Path ==“/” 的 URL。

如果一个子树已经被注册并且收到了一个请求，该请求在没有尾部斜杠的情况下命名子树根，ServeMux 将该请求重定向到子树根（添加尾部斜杠）。可以使用没有尾部斜杠的单独注册路径来覆盖此行为。例如，注册 “/images/” 会导致 ServeMux 将对 “/images” 的请求重定向到“/images/”，除非 “/images” 已单独注册。

模式可以选择以主机名开头，将匹配限制为仅在该主机上的 URL。特定于主机的模式优先于一般模式，因此处理程序可能会注册这两个模式 “/codesearch” 和 “codesearch.google.com/”，而不会同时接管对 “http://www.google.com/” 的请求”。

ServeMux 还负责清理 URL 请求路径和主机标头，剥离端口号并重定向任何包含 . 或 .. 元素或重复的斜杠到等效的、更清晰的 URL。

### func NewServeMux

```go

```

### func (*ServeMux) Handle

```go

```

### func (*ServeMux) HandleFunc

```go

```

### func (*ServeMux) Handler

```go

```

### func (*ServeMux) ServeHTTP

```go

```

### type Server

```go
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string

	Handler Handler // handler to invoke, http.DefaultServeMux if nil

	// DisableGeneralOptionsHandler, if true, passes "OPTIONS *" requests to the Handler,
	// otherwise responds with 200 OK and Content-Length: 0.
	DisableGeneralOptionsHandler bool

	// TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
	TLSConfig *tls.Config

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. A zero or negative value means
	// there will be no timeout.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body. If ReadHeaderTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	ReadHeaderTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	// A zero or negative value means there will be no timeout.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	IdleTimeout time.Duration

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes int

	// TLSNextProto optionally specifies a function to take over
	// ownership of the provided TLS connection when an ALPN
	// protocol upgrade has occurred. The map key is the protocol
	// name negotiated. The Handler argument should be used to
	// handle HTTP requests and will initialize the Request's TLS
	// and RemoteAddr if not already set. The connection is
	// automatically closed when the function returns.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	// ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	ConnState func(net.Conn, ConnState)

	// ErrorLog specifies an optional logger for errors accepting
	// connections, unexpected behavior from handlers, and
	// underlying FileSystem errors.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog *log.Logger

	// BaseContext optionally specifies a function that returns
	// the base context for incoming requests on this server.
	// The provided Listener is the specific Listener that's
	// about to start accepting requests.
	// If BaseContext is nil, the default is context.Background().
	// If non-nil, it must return a non-nil context.
	BaseContext func(net.Listener) context.Context

	// ConnContext optionally specifies a function that modifies
	// the context used for a new connection c. The provided ctx
	// is derived from the base context and has a ServerContextKey
	// value.
	ConnContext func(ctx context.Context, c net.Conn) context.Context
	// contains filtered or unexported fields
}
```

Server 定义运行 HTTP 服务器的参数。服务器的零值是有效配置。

### func (*Server) Close

```go

```

### func (*Server) ListenAndServe

```go

```

### func (*Server) ListenAndServeTLS

```go

```

### func (*Server) Serve

```go

```

### func (*Server) ServeTLS

```go

```

### func (*Server) SetKeepAlivesEnabled

```go

```

### func (*Server) Shutdown

```go

```

### type Transport

```go
type Transport struct {

	// Proxy specifies a function to return a proxy for a given
	// Request. If the function returns a non-nil error, the
	// request is aborted with the provided error.
	//
	// The proxy type is determined by the URL scheme. "http",
	// "https", and "socks5" are supported. If the scheme is empty,
	// "http" is assumed.
	//
	// If Proxy is nil or returns a nil *URL, no proxy is used.
	Proxy func(*Request) (*url.URL, error)

	// OnProxyConnectResponse is called when the Transport gets an HTTP response from
	// a proxy for a CONNECT request. It's called before the check for a 200 OK response.
	// If it returns an error, the request fails with that error.
	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error

	// DialContext specifies the dial function for creating unencrypted TCP connections.
	// If DialContext is nil (and the deprecated Dial below is also nil),
	// then the transport dials using package net.
	//
	// DialContext runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later DialContext completes.
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// Dial specifies the dial function for creating unencrypted TCP connections.
	//
	// Dial runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later Dial completes.
	//
	// Deprecated: Use DialContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialContext takes priority.
	Dial func(network, addr string) (net.Conn, error)

	// DialTLSContext specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// If DialTLSContext is nil (and the deprecated DialTLS below is also nil),
	// DialContext and TLSClientConfig are used.
	//
	// If DialTLSContext is set, the Dial and DialContext hooks are not used for HTTPS
	// requests and the TLSClientConfig and TLSHandshakeTimeout
	// are ignored. The returned net.Conn is assumed to already be
	// past the TLS handshake.
	DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// DialTLS specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// Deprecated: Use DialTLSContext instead, which allows the transport
	// to cancel dials as soon as they are no longer needed.
	// If both are set, DialTLSContext takes priority.
	DialTLS func(network, addr string) (net.Conn, error)

	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client.
	// If nil, the default configuration is used.
	// If non-nil, HTTP/2 support may not be enabled by default.
	TLSClientConfig *tls.Config

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration

	// DisableKeepAlives, if true, disables HTTP keep-alives and
	// will only use the connection to the server for a single
	// HTTP request.
	//
	// This is unrelated to the similarly named TCP keep-alives.
	DisableKeepAlives bool

	// DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression bool

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	//
	// Zero means no limit.
	MaxConnsPerHost int

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration

	// ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	ResponseHeaderTimeout time.Duration

	// ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout time.Duration

	// TLSNextProto specifies how the Transport switches to an
	// alternate protocol (such as HTTP/2) after a TLS ALPN
	// protocol negotiation. If Transport dials an TLS connection
	// with a non-empty protocol name and TLSNextProto contains a
	// map entry for that key (such as "h2"), then the func is
	// called with the request's authority (such as "example.com"
	// or "example.com:1234") and the TLS connection. The function
	// must return a RoundTripper that then handles the request.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper

	// ProxyConnectHeader optionally specifies headers to send to
	// proxies during CONNECT requests.
	// To set the header dynamically, see GetProxyConnectHeader.
	ProxyConnectHeader Header

	// GetProxyConnectHeader optionally specifies a func to return
	// headers to send to proxyURL during a CONNECT request to the
	// ip:port target.
	// If it returns an error, the Transport's RoundTrip fails with
	// that error. It can return (nil, nil) to not add headers.
	// If GetProxyConnectHeader is non-nil, ProxyConnectHeader is
	// ignored.
	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)

	// MaxResponseHeaderBytes specifies a limit on how many
	// response bytes are allowed in the server's response
	// header.
	//
	// Zero means to use a default limit.
	MaxResponseHeaderBytes int64

	// WriteBufferSize specifies the size of the write buffer used
	// when writing to the transport.
	// If zero, a default (currently 4KB) is used.
	WriteBufferSize int

	// ReadBufferSize specifies the size of the read buffer used
	// when reading from the transport.
	// If zero, a default (currently 4KB) is used.
	ReadBufferSize int

	// ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
	// Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
	// By default, use of any those fields conservatively disables HTTP/2.
	// To use a custom dialer or TLS config and still attempt HTTP/2
	// upgrades, set this to true.
	ForceAttemptHTTP2 bool
	// contains filtered or unexported fields
}
```

Transport 是 RoundTripper 的一个实现，它支持 HTTP、HTTPS 和 HTTP 代理（用于 HTTP 或带 CONNECT 的 HTTPS）。

默认情况下，传输缓存连接以供将来重用。当访问许多主机时，这可能会留下许多打开的连接。可以使用 Transport 的 CloseIdleConnections 方法以及 MaxIdleConnsPerHost 和 DisableKeepAlives 字段来管理此行为。

传输应该被重用而不是根据需要创建。传输对于多个 goroutines 的并发使用是安全的。

Transport 是用于发出 HTTP 和 HTTPS 请求的低级原语。有关高级功能，例如 cookie 和重定向，请参阅客户端。

Transport 对 HTTP URL 使用 HTTP/1.1，对 HTTPS URL 使用 HTTP/1.1 或 HTTP/2，具体取决于服务器是否支持 HTTP/2，以及 Transport 的配置方式。 DefaultTransport 支持 HTTP/2。要在传输上显式启用 HTTP/2，请使用 golang.org/x/net/http2 并调用 ConfigureTransport。有关 HTTP/2 的更多信息，请参阅包文档。

状态代码在 1xx 范围内的响应要么被自动处理（100 expect-continue）要么被忽略。一个例外是 HTTP 状态代码 101（切换协议），它被视为终端状态并由 RoundTrip 返回。要查看忽略的 1xx 响应，请使用 httptrace 跟踪包的 ClientTrace.Got1xxResponse。

如果请求是幂等的并且没有正文或定义了 Request.GetBody，则传输仅在遇到网络错误时重试请求。如果 HTTP 请求具有 HTTP 方法 GET、HEAD、OPTIONS 或 TRACE，则它们被认为是幂等的；或者如果他们的 Header 映射包含 “Idempotency-Key” 或 “X-Idempotency-Key” 条目。如果幂等键值是零长度切片，则请求被视为幂等但不会在线发送标头。

### func (*Transport) CancelRequest (DEPRECATED)

```go

```

### func (*Transport) Clone

```go

```

### func (*Transport) CloseIdleConnections

```go

```

### func (*Transport) RegisterProtocol

```go

```

### func (*Transport) RoundTrip

```go

```

## 目录

- [cgi](http/cgi.md)
- [cookiejar](http/cookiejar.md)
- [fcgi](http/fcgi.md)
- [httptest](http/httptest.md)
- [httptrace](http/httptrace.md)
- [httputil](http/httputil.md)
- [pprof](http/pprof.md)
- [internal](http/internal.md)
    - [ascii](http/internal/ascii.md)
    - [testcert](http/internal/testcert.md)
