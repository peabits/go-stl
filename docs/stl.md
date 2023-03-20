
# Standard library

## archive

### [tar](archive/tar.md)

tar包实现了对tar档案的访问。

### [zip](archive/zip.md)

zip包提供了对ZIP档案的读写支持。

## [bufio](bufio.md)

bufio包实现了缓冲的I/O。它包装了一个io.Reader或io.Writer对象，创建了另一个对象（Reader或Writer），它也实现了该接口，但提供了缓冲和一些文本I/O的帮助。

## [builtin](builtin.md)

builtin包为Go的预声明标识符提供了文档。

## [bytes](bytes.md)

bytes包实现了操作字节切片的功能。

## compress

### [bzip2](compress/bzip2.md)

### [flate](compress/flate.md)

### [gzip](compress/gzip.md)

### [lzw](compress/lzw.md)

### [zlib](compress/zlib.md)

## container

### [heap](container/heap.md)

heap包为任何实现heap.Interface的类型提供堆操作。

### [list](container/list.md)

list包实现了一个双链表。

### [ring](container/ring.md)

ring包实现了对循环列表的操作。

## [context](context.md)

context包定义了Context类型，它携带了截止日期、取消信号和其他跨API边界和进程间的请求范围的值。

## [crypto](crypto.md)

crypto包收集了常见的加密常数。

### [aes](crypto/aes.md)

### [cipher](crypto/cipher.md)

### [des](crypto/des.md)

### [dsa](crypto/dsa.md)

### [ecdh](crypto/ecdh.md)

### [ecdsa](crypto/ecdsa.md)

### [ed25519](crypto/ed25519.md)

### [elliptic](crypto/elliptic.md)

### [hmac](crypto/hmac.md)

### [md5](crypto/md5.md)

### [rand](crypto/rand.md)

### [rc4](crypto/rc4.md)

### [rsa](crypto/rsa.md)

### [sha1](crypto/sha1.md)

### [sha256](crypto/sha256.md)

### [sha512](crypto/sha512.md)

### [subtle](crypto/subtle.md)

### [tls](crypto/tls.md)

### [x509](crypto/x509.md)

### [x509/pkix](crypto/x509/pkix.md)

## database

### [sql](database/sql.md)

sql包提供了一个围绕SQL（或类似SQL）数据库的通用接口。

### [sql/driver](database/sql/driver.md)

driver包定义了数据库驱动要实现的接口，如包sql所使用。

## debug

## [embed](embed.md)

embed包提供了对嵌入运行中的Go程序的文件的访问。

## [encoding](encoding.md)

encoding包定义了其他包共享的接口，这些接口将数据转换为字节级和文本表示法。

## [errors](errors.md)

errors包实现了处理错误的函数。

## [expvar](expvar.md)

expvar包为公共变量提供了一个标准化的接口，例如服务器中的操作计数器。

## [flag](flog.md)

flag包实现了命令行flag的解析。

## [fmt](fmt.md)

fmt包用类似于C的printf和scanf的函数来实现格式化的I/O。

## go

## [hash](hash.md)

hash包提供了hash函数的接口。

## [html](html.md)

html包提供了用于转义和取消转义的HTML文本的函数。

### [template](html/template.md)

template包实现了数据驱动的模板，用于生成安全的HTML输出，防止代码注入。

## [image](image.md)

image包实现了一个基本的二维图像库。

## index

## [io](io.md)

io包提供了I/O原语的基本接口。

## [log](log.md)

log包实现了一个简单的日志包。

## [math](math.md)

mdth包提供基本常数和数学函数。

## [mime](mime.md)

mime包实现了MIME规范的部分内容

## [net](net.md)

net包为网络I/O提供了一个可移植的接口，包括TCP/IP、UDP、域名解析和Unix域套接字。

## [os](os.md)

os包提供了一个独立于平台的操作系统功能接口。

## [path](path.md)

path包实现了操作斜线分隔的路径的实用程序。

## [plugin](plugin.md)

plugin包实现了Go插件的加载和符号解析。

## [reflect](reflect.md)

reflect包实现了运行时反射，允许程序操作具有任意类型的对象。

## [regexp](regexp.md)

regexp包实现了正则表达式搜索。

## [runtime](runtime.md)

runtime包包含与Go的运行时系统互动的操作，例如控制goroutines的函数。

## [sort](sort.md)

sort包提供了对切片和用户定义的集合进行排序的基元。

## [strconv](strconv.md)

strconv包实现了基本数据类型的字符串表示法之间的转换。

## [strings](strings.md)

strings包实现了简单的函数来操作UTF-8编码的字符串。

## [sync](sync.md)

sync包提供了基本的同步基元，如互斥锁。

## [syscall](syscall.md)

syscall包包含了一个通往低级操作系统基元的接口。

## [testing](testing.md)

testing包为Go包的自动测试提供支持。

## text

## [time](time.md)

time包提供测量和显示时间的功能。

## [unicode](unicode.md)

unicode包提供数据和函数来测试Unicode码位的一些属性。

## [unsafe](unsafe.md)

unsafe包包含了围绕Go程序的类型安全的操作。

## internal
