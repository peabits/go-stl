
# Standard library

## archive

### [tar](archive/tar.md)

tar 包实现了对 tar 档案的访问。

### [zip](archive/zip.md)

zip 包提供对读取和写入 ZIP 档案的支持。

## [bufio](bufio.md)

bufio 包实现缓冲 I/O。它包装了一个 io.Reader 或 io.Writer 对象，创建了另一个对象（Reader 或 Writer），该对象也实现了该接口，但为文本 I/O 提供了缓冲和一些帮助。

## [builtin](builtin.md)

builtin 包为 Go 的预声明标识符提供了文档。

## [bytes](bytes.md)

bytes 包实现了用于操作字节片的函数。

## compress

### [bzip2](compress/bzip2.md)

bzip2 包实现 bzip2 解压缩。

### [flate](compress/flate.md)

flate 包实现了 DEFLATE 压缩数据格式，如 RFC 1951 中所述。

### [gzip](compress/gzip.md)

gzip 包实现了 gzip 格式压缩文件的读写，如 RFC 1952 中所规定。

### [lzw](compress/lzw.md)

lzw 包实现 Lempel-Ziv-Welch 压缩数据格式，在 T. A. Welch，“高性能数据压缩技术”，计算机，17(6)（1984 年 6 月），第 8-19 页中有描述。

### [zlib](compress/zlib.md)

zlib 包实现了 zlib 格式压缩数据的读写，如 RFC 1950 中所指定。

## container

### [heap](container/heap.md)

heap 包为任何实现 heap.Interface 的类型提供堆操作。

### [list](container/list.md)

list 包实现了双向链表。

### [ring](container/ring.md)

ring 包实现对循环列表的操作。

## [context](context.md)

context 包定义了上下文类型，它携带截止日期、取消信号和其他跨 API 边界和进程之间的请求范围的值。

## [crypto](crypto.md)

crypto 包收集常见的密码常量。

### [aes](crypto/aes.md)

aes 包实现 AES 加密（以前称为 Rijndael），如美国联邦信息处理标准出版物 197 中所定义。

### [cipher](crypto/cipher.md)

cipher 包实现标准分组密码模式，可以围绕低级分组密码实现进行包装。

### [des](crypto/des.md)

des 包实现了美国联邦信息处理标准出版物 46-3 中定义的数据加密标准 (DES) 和三重数据加密算法 (TDEA)。

### [dsa](crypto/dsa.md)

dsa 包实现了 FIPS 186-3 中定义的数字签名算法。

### [ecdh](crypto/ecdh.md)

ecdh 包在 NIST 曲线和 Curve25519 上实现了椭圆曲线 Diffie-Hellman。

### [ecdsa](crypto/ecdsa.md)

ecdsa 包实现了椭圆曲线数字签名算法，如 FIPS 186-4 和 SEC 1，版本 2.0 中所定义。

### [ed25519](crypto/ed25519.md)

ed25519 包实现了 Ed25519 签名算法。

### [elliptic](crypto/elliptic.md)

elliptic 包在素数域上实现了标准的 NIST P-224、P-256、P-384 和 P-521 椭圆曲线。

### [hmac](crypto/hmac.md)

hmac 包实现了美国联邦信息处理标准出版物 198 中定义的密钥散列消息身份验证代码 (HMAC)。

### [md5](crypto/md5.md)

md5 包实现了 RFC 1321 中定义的 MD5 散列算法。

### [rand](crypto/rand.md)

rand 包实现了一个加密安全的随机数生成器。

### [rc4](crypto/rc4.md)

rc4 包实现了 RC4 加密，如 Bruce Schneier 的 Applied Cryptography 中所定义。

### [rsa](crypto/rsa.md)

rsa 包实现了 PKCS #1 和 RFC 8017 中指定的 RSA 加密。

### [sha1](crypto/sha1.md)

sha1 包实现了 RFC 3174 中定义的 SHA-1 散列算法。

### [sha256](crypto/sha256.md)

sha256 包实现了 FIPS 180-4 中定义的 SHA224 和 SHA256 哈希算法。

### [sha512](crypto/sha512.md)

sha512 包实现了 FIPS 180-4 中定义的 SHA-384、SHA-512、SHA-512/224 和 SHA-512/256 哈希算法。

### [subtle](crypto/subtle.md)

subtle 包实现了通常在密码代码中有用但需要仔细考虑才能正确使用的功能。

### [tls](crypto/tls.md)

tls 包部分实现了 RFC 5246 中指定的 TLS 1.2 和 RFC 8446 中指定的 TLS 1.3。

### [x509](crypto/x509.md)

x509 包实现了 X.509 标准的一个子集。

### [x509/pkix](crypto/x509/pkix.md)

pkix 包包含用于 ASN.1 解析和序列化 X.509 证书、CRL 和 OCSP 的共享低级结构。

## database

### [sql](database/sql.md)

sql 包提供了一个围绕 SQL（或类似 SQL）数据库的通用接口。

### [sql/driver](database/sql/driver.md)

driver 包程序定义了由 sql 包使用的数据库驱动程序实现的接口。

## debug

### [buildinfo](debug/buildinfo.md)

buildinfo 包提供了嵌入了 GO 二进制中有关其构建方式的信息。

### [dwarf](debug/dwarf.md)

dwarf 包提供对从可执行文件加载的 DWARF 调试信息的访问，如 DWARF 2.0 标准中所定义，网址为 http://dwarfstd.org/doc/dwarf-2.0.0.pdf。

### [elf](debug/elf.md)

elf 包实现了对 ELF 对象文件的访问。

### [gosym](debug/gosym.md)

gosym 包实现了对 gc 编译器生成的 Go 二进制文件中嵌入的 Go 符号和行号表的访问。

### [macho](debug/macho.md)

macho 包实现了对 Mach-O 目标文件的访问。

### [pe](debug/pe.md)

pe 包实现对 PE（Microsoft Windows 可移植可执行文件）文件的访问。

### [plan9obj](debug/plan9obj.md)

plan9obj 包实现了对 Plan 9 a.out 目标文件的访问。

## [embed](embed.md)

embed 包提供对嵌入在运行的 Go 程序中的文件的访问。

## [encoding](encoding.md)

encoding 包定义了其他包共享的接口，这些包将数据与字节级和文本表示形式相互转换。

### [ascii85](encoding/ascii85.md)

ascii85 包实现了 btoa 工具和 Adob​​e 的 PostScript 和 PDF 文档格式中使用的 ascii85 数据编码。

### [asn1](encoding/asn1.md)

asn1 包实现 DER 编码的 ASN.1 数据结构的解析，如 ITU-T Rec X.690 中所定义。

### [base32](encoding/base32.md)

base32 包实现了 RFC 4648 指定的 base32 编码。

### [base64](encoding/base64.md)

base64 包实现了 RFC 4648 指定的 base64 编码。

### [binary](encoding/binary.md)

binary 包实现了数字和字节序列之间的简单转换以及 varint 的编码和解码。

### [csv](encoding/csv.md)

csv 包读取和写入逗号分隔值 (CSV) 文件。

### [gob](encoding/gob.md)

gob 包管理 gob 流 - 在编码器（发送器）和解码器（接收器）之间交换的二进制值。

### [hex](encoding/hex.md)

hex 包实现了十六进制编码和解码。

### [json](encoding/json.md)

json 包实现了 RFC 7159 中定义的 JSON 编码和解码。

### [pem](encoding/pem.md)

pem 包实现了 PEM 数据编码，起源于 Privacy Enhanced Mail。

### [xml](encoding/xml.md)

xml 包实现了一个简单的 XML 1.0 解析器，它可以理解 XML 名称空间。

## [errors](errors.md)

errors 包实现了操作错误的功能。

## [expvar](expvar.md)

expvar 包为公共变量提供标准化接口，例如服务器中的操作计数器。

## [flag](flag.md)

flag 包实现命令行标志解析。

## [fmt](fmt.md)

fmt 包使用类似于 C 的 printf 和 scanf 的函数来实现格式化 I/O。

## go

### [ast](go/ast.md)

ast 包声明用于表示 Go 包语法树的类型。

### [build](go/build.md)

build 包收集有关 Go 包的信息。

### [build/constraint](go/build/constraint.md)

constraint 包实现构建约束线的解析和评估。

### [constant](go/constant.md)

constant 包实现了表示无类型 Go 常量及其相应操作的值。

### [doc](go/doc.md)

doc 包从 Go AST 中提取源代码文档。

### [doc/comment](go/doc/comment.md)

comment 包实现 Go 文档注释（文档注释）的解析和重新格式化，这些注释紧接在 package、const、func、type 或 var 的顶级声明之前。

### [format](go/format.md)

format 包实现了 Go 源码的标准格式化。

### [importer](go/importer.md)

importer 包提供对导出数据导入器的访问。

### [parser](go/parser.md)

parser 包实现了 Go 源文件的解析器。

### [printer](go/printer.md)

printer 包实现 AST 节点的打印。

### [scanner](go/scanner.md)

scanner 包实现了 Go 源文本的扫描器。

### [token](go/token.md)

token 包定义了表示 Go 编程语言的词法令牌和令牌上的基本操作（打印、谓词）的常量。

### [types](go/types.md)

types 包声明数据类型并实现 Go 包类型检查的算法。

## [hash](hash.md)

hash 包提供了哈希函数的接口。

### [adler32](hash/adler32.md)

adler32 包实现了 Adler-32 校验和。

### [crc32](hash/crc32.md)

crc32 包实现了 32 位循环冗余校验或 CRC-32 校验和。

### [crc64](hash/crc64.md)

crc64 包实现了 64 位循环冗余校验或 CRC-64 校验和。

### [fnv](hash/fnv.md)

fnv 包实现了 FNV-1 和 FNV-1a，这是由 Glenn Fowler、Landon Curt Noll 和 Phong Vo 创建的非加密散列函数。

### [maphash](hash/maphash.md)

maphash 包提供字节序列的散列函数。

## [html](html.md)

html 包提供了转义和取消转义 HTML 文本的功能。

### [template](html/template.md)

template 包实现了数据驱动的模板，用于生成可防止代码注入的 HTML 输出。

## [image](image.md)

image 包实现了一个基本的二维图像库。

### [color](image/color.md)

color 包实现了一个基本的颜色库。

### [color/palette](image/color/palette.md)

palette 包提供了标准的调色板。

### [draw](image/draw.md)

draw 包提供图像合成功能。

### [gif](image/gif.md)

gif 包实现了 GIF 图像解码器和编码器。

### [jpeg](image/jpeg.md)

jpeg 包实现了 JPEG 图像解码器和编码器。

### [png](image/png.md)

png 包实现了 PNG 图像解码器和编码器。

## index

### [suffixarray](index/suffixarray.md)

suffixarray 包使用内存后缀数组以对数时间实现子字符串搜索。

## [io](io.md)

io 包提供了 I/O 原语的基本接口。

### [fs](io/fs.md)

fs 包定义了文件系统的基本接口。

### [ioutil](io/ioutil.md)

ioutil 包实现了一些 I/O 实用函数。

## [log](log.md)

log 包实现了一个简单的日志记录包。

### [syslog](log/syslog.md)

syslog 包为系统日志服务提供了一个简单的接口。

## [math](math.md)

math 包提供基本常量和数学函数。

### [big](math/big.md)

big 包实现任意精度算术（大数）。

### [bits](math/bits.md)

bits 包为预先声明的无符号整数类型实现位计数和操作功能。

### [cmplx](math/cmplx.md)

cmplx 包为复数提供了基本常量和数学函数。

### [rand](math/rand.md)

rand 包实现了不适合安全敏感工作的伪随机数生成器。

## [mime](mime.md)

mime 包实现了部分 MIME 规范。

### [multipart](mime/multipart.md)

multipart 包实现 MIME 多部分解析，如 RFC 2046 中所定义。

### [quotedprintable](mime/quotedprintable.md)

quotedprintable 包实现了 RFC 2045 指定的 quoted-printable 编码。

## [net](net.md)

net 包为网络 I/O 提供了一个可移植的接口，包括 TCP/IP、UDP、域名解析和 Unix 域套接字。

### [http](net/http.md)

http 包提供 HTTP 客户端和服务器实现。

### [http/cgi](net/http/cgi.md)

cgi 包实现了 RFC 3875 中指定的 CGI（通用网关接口）。

### [http/cookiejar](net/http/cookiejar.md)

cookiejar 包实现了内存中符合 RFC 6265 标准的 http.CookieJar。

### [http/fcgi](net/http/fcgi.md)

fcgi 包实现了 FastCGI 协议。

### [http/httptest](net/http/httptest.md)

httptest 包提供用于 HTTP 测试的实用程序。

### [http/httptrace](net/http/httptrace.md)

httptrace 包提供了跟踪 HTTP 客户端请求中的事件的机制。

### [http/httputil](net/http/httputil.md)

httputil 包提供 HTTP 实用程序功能，补充了 net/http 包中更常见的功能。

### [http/pprof](net/http/pprof.md)

pprof 包通过其 HTTP 服务器运行时分析数据以 pprof 可视化工具预期的格式提供服务。

### [mail](net/mail.md)

mail 包实现邮件消息的解析。

### [netip](net/netip.md)

netip 包定义了一个 IP 地址类型，它是一个小值类型。

### [rpc](net/rpc.md)

rpc 包提供通过网络或其他 I/O 连接访问对象的导出方法。

### [rpc/jsonrpc](net/rpc/jsonrpc.md)

jsonrpc 包为 rpc 包实现了 JSON-RPC 1.0 ClientCodec 和 ServerCodec。

### [smtp](net/smtp.md)

smtp 包实现了 RFC 5321 中定义的简单邮件传输协议。

### [textproto](net/textproto.md)

textproto 包实现了对 HTTP、NNTP 和 SMTP 样式的基于文本的请求/响应协议的通用支持。

### [url](net/url.md)

url 包解析 URL 并实现查询转义。

## [os](os.md)

os 为操作系统功能提供了一个独立于平台的接口。

### [exec](os/exec.md)

exec 包运行外部命令。

### [signal](os/signal.md)

signal 包实现对传入信号的访问。

### [user](os/user.md)

user 包允许按名称或 ID 查找用户帐户。

## [path](path.md)

path 包实现用于操作斜杠分隔路径的实用程序。

### [filepath](path/filepath.md)

filepath 包实现实用例程，以与目标操作系统定义的文件路径兼容的方式操作文件名路径。

## [plugin](plugin.md)

plugin 包实现了 Go 插件的加载和符号解析。

## [reflect](reflect.md)

reflect 包实现运行时反射，允许程序操作任意类型的对象。

## [regexp](regexp.md)

regexp 包实现正则表达式搜索。

### [syntax](regexp/syntax.md)

syntax 包将正则表达式解析成解析树，并将解析树编译成程序。

## [runtime](runtime.md)

runtime 包包含与 Go 的运行时系统交互的操作，例如控制 goroutines 的函数。

### [cgo](runtime/cgo.md)

cgo 包包含对 cgo 工具生成的代码的运行时支持。

### [coverage](runtime/coverage.md)

### [debug](runtime/debug.md)

debug 包包含程序在运行时进行自我调试的功能。

### [metrics](runtime/metrics.md)

metrics 包提供了一个稳定的接口来访问由 Go 运行时导出的实现定义的指标。

### [pprof](runtime/pprof.md)

pprof 包以 pprof 可视化工具期望的格式写入运行时分析数据。

### [race](runtime/race.md)

race 包实现数据竞争检测逻辑。

### [trace](runtime/trace.md)

trace 包包含程序为 Go 执行跟踪器生成跟踪的工具。

## [sort](sort.md)

sort 包提供了用于排序切片和用户定义的集合的原语。

## [strconv](strconv.md)

strconv 包实现了与基本数据类型的字符串表示之间的转换。

## [strings](strings.md)

strings 包实现了简单的函数来操作 UTF-8 编码的字符串。

## [sync](sync.md)

sync 包提供基本的同步原语，例如互斥锁。

### [atomic](sync/atomic.md)

atomic 包提供了可用于实现同步算法的低级原子内存原语。

## [syscall](syscall.md)

syscall 包包含低级操作系统原语的接口。

### [js](syscall/js.md)

js 包可以访问 WebAssembly 主机环境（当使用 js/wasm 架构时）。

## [testing](testing.md)

testing 包为 Go 包的自动化测试提供支持。

### [fstest](testing/fstest.md)

fstest 包实现了对文件系统的测试实现和用户的支持。

### [iotest](testing/iotest.md)

iotest 包实现了主要用于测试的 Reader 和 Writers。

### [quick](testing/quick.md)

quick 包实现实用功能以帮助进行黑盒测试。

## text

### [scanner](text/scanner.md)

scanner 包为 UTF-8 编码的文本提供扫描器和分词器。

### [tabwriter](text/tabwriter.md)

tabwriter 包实现了一个写入过滤器 (tabwriter.Writer)，它将输入中的选项卡式列转换为正确对齐的文本。

### [template](text/template.md)

template 包实现了用于生成文本输出的数据驱动模板。

### [template/parse](text/template/parse.md)

parse 包为由 text/template 和 html/template 定义的模板构建解析树。

## [time](time.md)

time 包提供了测量和显示时间的功能。

### [tzdata](time/tzdata.md)

tzdata 包提供了时区数据库的嵌入式副本。

## [unicode](unicode.md)

unicode 包提供数据和函数来测试 Unicode 代码点的某些属性。

### [utf16](unicode/utf16.md)

utf16 包实现了 UTF-16 序列的编码和解码。

### [utf8](unicode/utf8.md)

utf8 包实现函数和常量以支持以 UTF-8 编码的文本。

## [unsafe](unsafe.md)

unsafe 包包含绕过 Go 程序类型安全的操作。

## internal

### [abi](internal/abi.md)

### [buildcfg](internal/buildcfg.md)

buildcfg 包提供对当前环境描述的构建配置的访问。

### [bytealg](internal/bytealg.md)

### [cfg](internal/cfg.md)

cfg 包包含由 Go 命令和 internal/testen 共享的配置

### [coverage](internal/coverage.md)

### [coverage/calloc](internal/coverage/calloc.md)

### [coverage/cformat](internal/coverage/cformat.md)

### [coverage/cmerge](internal/coverage/cmerge.md)

### [coverage/decodecounter](internal/coverage/decodecounter.md)

### [coverage/decodemeta](internal/coverage/decodemeta.md)

### [coverage/encodecounter](internal/coverage/encodecounter.md)

### [coverage/encodemeta](internal/coverage/encodemeta.md)

### [coverage/pods](internal/coverage/pods.md)

### [coverage/rtcov](internal/coverage/rtcov.md)

### [coverage/slicereader](internal/coverage/slicereader.md)

### [coverage/slicewriter](internal/coverage/slicewriter.md)

### [coverage/stringtab](internal/coverage/stringtab.md)

### [coverage/uleb128](internal/coverage/uleb128.md)

### [cpu](internal/cpu.md)

cpu 包实现了 Go 标准库使用的处理器特征检测。

### [dag](internal/dag.md)

dag 包实现了一种用于表达有向无环图的语言。

### [diff](internal/diff.md)

### [fmtsort](internal/fmtsort.md)

fmtsort 包代表 fmt 和 text/template 包为地图提供了一个通用的稳定排序机制。

### [fuzz](internal/fuzz.md)

fuzz 包为使用 “go test” 构建的测试和在测试包中使用模糊测试功能的程序提供了通用的模糊测试功能。

### [goarch](internal/goarch.md)

goarch 包包含 GOARCH 特定的常量。

### [godebug](internal/godebug.md)

godebug 包使 $GODEBUG 环境变量中的设置可用于其他包。

### [goexperiment](internal/goexperiment.md)

goexperiment 包实现了对工具链实验的支持。

### [goos](internal/goos.md)

goos 包包含特定于 GOOS 的常量。

### [goroot](internal/goroot.md)

### [goversion](internal/goversion.md)

### [intern](internal/intern.md)

intern 包允许您通过将较大的可比较值（例如 16 字节字符串标头）装箱到全局唯一的 8 字节指针中来生成较小的可比较值。

### [itoa](internal/itoa.md)

### [lazyregexp](internal/lazyregexp.md)

lazyregexp 包是 regexp 的薄包装器，允许使用全局 regexp 变量而不强制它们在 init 时编译。

### [lazytemplate](internal/lazytemplate.md)

lazytemplate 包是文本/模板的薄包装器，允许使用全局模板变量而不强制它们在初始化时被解析。

### [nettrace](internal/nettrace.md)

nettrace 包包含用于跟踪 net 包中活动的内部挂钩。

### [obscuretestdata](internal/obscuretestdata.md)

obscuretestdata 包包含测试使用的功能，以便更轻松地处理主要由于 golang.org/issue/34986 必须被隐藏的测试数据。

### [oserror](internal/oserror.md)

oserror 包定义了 os 包中使用的错误值。

### [pkgbits](internal/pkgbits.md)

pkgbits 包为 Unified IR 的导出数据格式实现低级编码抽象。

### [platform](internal/platform.md)

### [poll](internal/poll.md)

poll 包支持对文件描述符进行轮询的非阻塞 I/O。

### [profile](internal/profile.md)

profile 包提供了 github.com/google/pprof/proto/profile.proto 的表示以及以这种格式编码/解码/合并配置文件的方法。

### [race](internal/race.md)

race 包包含用于手动检测竞争检测器代码的辅助函数。

### [reflectlite](internal/reflectlite.md)

reflectlite 包实现了 reflect 的轻量级版本，除了“runtime”和“unsafe”之外不使用任何包。

### [safefilepath](internal/safefilepath.md)

safefilepath 包操纵操作系统文件路径。

### [saferio](internal/saferio.md)

saferio 包提供了 I/O 功能，可以避免不必要地分配大量内存。

### [singleflight](internal/singleflight.md)

singleflight 包提供了重复函数调用抑制机制。

### [syscall/execenv](internal/syscall/execenv.md)

### [syscall/unix](internal/syscall/unix.md)

### [syscall/windows](internal/syscall/windows.md)

### [syscall/windows/registry](internal/syscall/windows/registry.md)

registry 包提供对 Windows 注册表的访问。

### [syscall/windows/sysdll](internal/syscall/windows/sysdll.md)

sysdll 包是一个内部叶包，它记录和报告 Go 本身使用了哪些 Windows DLL 名称。

### [sysinfo](internal/sysinfo.md)

sysinfo 包实现了可用于调试或信息目的的高级硬件信息收集。

### [testenv](internal/testenv.md)

testenv 包提供了有关在 Go 团队运行的不同测试环境中可用的功能的信息。

### [testlog](internal/testlog.md)

testlog 包提供了测试和包 os 之间的反向通道通信路径，以便 cmd/go 可以看到测试参考了哪些环境变量和文件。

### [testpty](internal/testpty.md)

testpty 包是一个简单的 Unix 系统伪终端包，通过 cgo 调用 C 函数实现。

### [trace](internal/trace.md)

### [txtar](internal/txtar.md)

txtar 包实现了一种简单的基于文本的文件存档格式。

### [types/errors](internal/types/errors.md)

### [unsafeheader](internal/unsafeheader.md)

unsafeheader 包包含 Go 运行时的切片和字符串实现的标头声明。

### [xcoff](internal/xcoff.md)

xcoff 包实现了对 XCOFF（扩展通用对象文件格式）文件的访问。
