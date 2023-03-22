
# os

## 概述

os 包为操作系统功能提供了一个独立于平台的接口。设计是类 Unix 的，尽管错误处理是类 Go 的；失败的调用返回错误类型的值而不是错误号。通常，错误中包含更多信息。例如，如果采用文件名的调用失败，如 Open 或 Stat，则错误将在打印时包含失败的文件名，并且类型为 *PathError，可以解包以获取更多信息。

os 接口旨在在所有操作系统中保持一致。通常不可用的功能出现在系统特定的包系统调用中。

这是一个简单的例子，打开一个文件并读取其中的一些内容。

```go
file, err := os.Open("file.go")  // For read access
if err != nil {
    log.Fatal(err)
}
```

如果打开失败，错误字符串将是不言自明的，比如

```go
open file.go: no such file or directory
```

然后可以将文件的数据读入字节片中。 Read 和 Write 从参数 sli 的长度中获取字节数

```go
data := make([]yte, 100)
count, err := file.Read(data)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
```

注意：对文件的最大并发操作数可能受操作系统或系统限制。该数字应该很高，但超过它可能会降低性能或导致其他问题。

## 索引

- [Constants](#常量)
- [Variables](#变量)
- [Functions](#函数)
    - [func Chdir(dir string) error](#func-chdir)
    - [func Chmod(name string, mode FileMode) error](#func-chmod)
    - [func Chown(name string, uid, gid int) error](#func-chown)
    - [func Chtimes(name string, atime time.Time, mtime time.Time) error](#func-chtimes)
    - [func Clearenv()](#func-clearenv)
    - [func DirFS(dir string) fs.FS](#func-dirfs)
    - [func Environ() []string](#func-environ)
    - [func Executable() (string, error)](#func-executable)
    - [func Exit(code int)](#func-exit)
    - [func Expand(s string, mapping func(string) string) string](#func-expand)
    - [func ExpandEnv(s string) string](#func-expandenv)
    - [func Getegid() int](#func-getegid)
    - [func Getenv(key string) string](#func-getenv)
    - [func Geteuid() int](#func-geteuid)
    - [func Getgid() int](#func-getgid)
    - [func Getgroups() ([]int, error)](#func-getgroups)
    - [func Getpagesize() int](#func-getpagesize)
    - [func Getpid() int](#func-getpid)
    - [func Getppid() int](#func-getppid)
    - [func Getuid() int](#func-getuid)
    - [func Getwd() (dir string, err error)](#func-getwd)
    - [func Hostname() (name string, err error)](#func-hostname)
    - [func IsExist(err error) bool](#func-isexist)
    - [func IsNotExist(err error) bool](#func-isnotexist)
    - [func IsPathSeparator(c uint8) bool](#func-ispathseparator)
    - [func IsPermission(err error) bool](#func-ispermission)
    - [func IsTimeout(err error) bool](#func-istimeout)
    - [func Lchown(name string, uid, gid int) error](#func-lchown)
    - [func Link(oldname, newname string) error](#func-link)
    - [func LookupEnv(key string) (string, bool)](#func-lookupenv)
    - [func Mkdir(name string, perm FileMode) error](#func-mkdir)
    - [func MkdirAll(path string, perm FileMode) error](#func-mkdirall)
    - [func MkdirTemp(dir, pattern string) (string, error)](#func-mkdirtemp)
    - [func NewSyscallError(syscall string, err error) error](#func-newsyscallerror)
    - [func Pipe() (r *File, w *File, err error)](#func-pipe)
    - [func ReadFile(name string) ([]byte, error)](#func-readfile)
    - [func Readlink(name string) (string, error)](#func-readlink)
    - [func Remove(name string) error](#func-remove)
    - [func RemoveAll(path string) error](#func-removeall)
    - [func Rename(oldpath, newpath string) error](#func-rename)
    - [func SameFile(fi1, fi2 FileInfo) bool](#func-samefile)
    - [func Setenv(key, value string) error](#func-setenv)
    - [func Symlink(oldname, newname string) error](#func-symlink)
    - [func TempDir() string](#func-tempdir)
    - [func Truncate(name string, size int64) error](#func-truncate)
    - [func Unsetenv(key string) error](#func-unsetenv)
    - [func UserCacheDir() (string, error)](#func-usercachedir)
    - [func UserConfigDir() (string, error)](#func-userconfigdir)
    - [func UserHomeDir() (string, error)](#func-userhomedir)
    - [func WriteFile(name string, data []byte, perm FileMode) error](#func-writefile)
- [Types](#类型)
    - [type DirEntry](#type-direntry)
        - [func ReadDir(name string) ([]DirEntry, error)](#func-readdir)
    - [type File](#type-file)
        - [func Create(name string) (*File, error)](#func-create)
        - [func CreateTemp(dir, pattern string) (*File, error)](#func-createtemp)
        - [func NewFile(fd uintptr, name string) *File](#func-newfile)
        - [func Open(name string) (*File, error)](#func-open)
        - [func OpenFile(name string, flag int, perm FileMode) (*File, error)](#func-openfile)
        - [func (f *File) Chdir() error](#func-file-chdir)
        - [func (f *File) Chmod(mode FileMode) error](#func-file-chmod)
        - [func (f *File) Chown(uid, gid int) error](#func-file-chown)
        - [func (f *File) Close() error](#func-file-close)
        - [func (f *File) Fd() uintptr](#func-file-fd)
        - [func (f *File) Name() string](#func-file-name)
        - [func (f *File) Read(b []byte) (n int, err error)](#func-file-read)
        - [func (f *File) ReadAt(b []byte, off int64) (n int, err error)](#func-file-readat)
        - [func (f *File) ReadDir(n int) ([]DirEntry, error)](#func-file-readdir)
        - [func (f *File) ReadFrom(r io.Reader) (n int64, err error)](#func-file-readfrom)
        - [func (f *File) Readdir(n int) ([]FileInfo, error)](#func-file-readdir-1)
        - [func (f *File) Readdirnames(n int) (names []string, err error)](#func-file-readdirnames)
        - [func (f *File) Seek(offset int64, whence int) (ret int64, err error)](#func-file-seek)
        - [func (f *File) SetDeadline(t time.Time) error](#func-file-setdeadline)
        - [func (f *File) SetReadDeadline(t time.Time) error](#func-file-setreaddeadline)
        - [func (f *File) SetWriteDeadline(t time.Time) error](#func-file-setwritedeadline)
        - [func (f *File) Stat() (FileInfo, error)](#func-file-stat)
        - [func (f *File) Sync() error](#func-file-sync)
        - [func (f *File) SyscallConn() (syscall.RawConn, error)](#func-file-syscallconn)
        - [func (f *File) Truncate(size int64) error](#func-file-truncate)
        - [func (f *File) Write(b []byte) (n int, err error)](#func-file-write)
        - [func (f *File) WriteAt(b []byte, off int64) (n int, err error)](#func-file-writeat)
        - [func (f *File) WriteString(s string) (n int, err error)](#func-file-writestring)
    - [type FileInfo](#type-fileinfo)
        - [func Lstat(name string) (FileInfo, error)](#func-lstat)
        - [func Stat(name string) (FileInfo, error)](#func-stat)
    - [type FileMode](#type-filemode)
    - [type LinkError](#type-linkerror)
        - [func (e *LinkError) Error() string](#func-linkerror-error)
        - [func (e *LinkError) Unwrap() error](#func-linkerror-unwrap)
    - [type PathError](#type-patherror)
    - [type ProcAttr](#type-procattr)
    - [type Process](#type-process)
        - [func FindProcess(pid int) (*Process, error)](#func-findprocess)
        - [func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)](#func-startprocess)
        - [func (p *Process) Kill() error](#func-process-kill)
        - [func (p *Process) Release() error](#func-process-release)
        - [func (p *Process) Signal(sig Signal) error](#func-process-signal)
        - [func (p *Process) Wait() (*ProcessState, error)](#func-process-wait)
    - [type ProcessState](#type-processstate)
        - [func (p *ProcessState) ExitCode() int](#func-processstate-exitcode)
        - [func (p *ProcessState) Exited() bool](#func-processstate-exited)
        - [func (p *ProcessState) Pid() int](#func-processstate-pid)
        - [func (p *ProcessState) String() string](#func-processstate-string)
        - [func (p *ProcessState) Success() bool](#func-processstate-success)
        - [func (p *ProcessState) Sys() any](#func-processstate-sys)
        - [func (p *ProcessState) SysUsage() any](#func-processstate-sysusage)
        - [func (p *ProcessState) SystemTime() time.Duration](#func-processstate-systemtime)
        - [func (p *ProcessState) UserTime() time.Duration](#func-processstate-usertime)
    - [type Signal](#type-signal)
    - [type SyscallError](#type-syscallerror)
        - [func (e *SyscallError) Error() string](#func-syscallerror-error)
        - [func (e *SyscallError) Timeout() bool](#func-syscallerror-timeout)
        - [func (e *SyscallError) Unwrap() error](#func-syscallerror-unwrap)

## 常量

```go
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
```

OpenFile 的标志包装了底层系统的标志。并非所有标志都可以在给定系统上实现。

```go
const (
    SEEK_SET int = 0 // seek relative to the origin of the file
	SEEK_CUR int = 1 // seek relative to the current offset
	SEEK_END int = 2 // seek relative to the end
)
```

求值从何而来。

弃用：使用 io.SeekStart、io.SeekCurrent 和 io.SeekEnd。

```go
const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
)
```

```go
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        = fs.ModeDir        // d: is a directory
	ModeAppend     = fs.ModeAppend     // a: append-only
	ModeExclusive  = fs.ModeExclusive  // l: exclusive use
	ModeTemporary  = fs.ModeTemporary  // T: temporary file; Plan 9 only
	ModeSymlink    = fs.ModeSymlink    // L: symbolic link
	ModeDevice     = fs.ModeDevice     // D: device file
	ModeNamedPipe  = fs.ModeNamedPipe  // p: named pipe (FIFO)
	ModeSocket     = fs.ModeSocket     // S: Unix domain socket
	ModeSetuid     = fs.ModeSetuid     // u: setuid
	ModeSetgid     = fs.ModeSetgid     // g: setgid
	ModeCharDevice = fs.ModeCharDevice // c: Unix character device, when ModeDevice is set
	ModeSticky     = fs.ModeSticky     // t: sticky
	ModeIrregular  = fs.ModeIrregular  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm // Unix permission bits, 0o777
)
```

定义的文件模式位是 FileMode 的最高有效位。九个最低有效位是标准的 Unix rwxrwxrwx 权限。这些位的值应被视为公共 API 的一部分，并可用于有线协议或磁盘表示：它们不得更改，但可能会添加新位。

## 变量

```go
var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = fs.ErrInvalid // "invalid argument"

	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"

	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
)
```

一些常见系统调用错误的可移植类比。

可以使用 errors.Is 针对这些错误测试从此包返回的错误。

```go
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

Stdin、Stdout 和 Stderr 是指向标准输入、标准输出和标准错误文件描述符的打开文件。

请注意，Go 运行时会针对恐慌和崩溃写入标准错误；关闭 Stderr 可能会导致这些消息转到其他地方，可能会转到稍后打开的文件。

```go
var Args []string
```

Args 包含命令行参数，以程序名称开头。

```go
var ErrProcessDone = errors.New("os: process already finished")
```

ErrProcessDone 指示进程已完成。

## 函数

### func Chdir

```go

```

### func Chmod

```go

```

### func Chown

```go

```

### func Chtimes

```go

```

### func Clearenv

```go

```

### func DirFS

```go

```

### func Environ

```go

```

### func Executable

```go

```

### func Exit

```go

```

### func Expand

```go

```

### func ExpandEnv

```go

```

### func Getegid

```go

```

### func Getenv

```go

```

### func Geteuid

```go

```

### func Getgid

```go

```

### func Getgroups

```go

```

### func Getpagesize

```go

```

### func Getpid

```go

```

### func Getppid

```go

```

### func Getuid

```go

```

### func Getwd

```go

```

### func Hostname

```go

```

### func IsExist

```go

```

### func IsNotExist

```go

```

### func IsPathSeparator

```go

```

### func IsPermission

```go

```

### func IsTimeout

```go

```

### func Lchown

```go

```

### func Link

```go

```

### func LookupEnv

```go

```

### func Mkdir

```go

```

### func MkdirAll

```go

```

### func MkdirTemp 

```go

```

### func NewSyscallError

```go

```

### func Pipe

```go

```

### func ReadFile

```go

```

### func Readlink

```go

```

### func Remove

```go

```

### func RemoveAll

```go

```

### func Rename

```go

```

### func SameFile

```go

```

### func Setenv

```go

```

### func Symlink

```go

```

### func TempDir

```go

```

### func Truncate

```go

```

### func Unsetenv

```go

```

### func UserCacheDir

```go

```

### func UserConfigDir

```go

```

### func UserHomeDir

```go

```

### func WriteFile

```go

```

## 类型

### type DirEntry

```go
type DirEntry = fs.DirEntry
```

### func ReadDir

```go

```

### type File

```go
type File struct {
	// contains filtered or unexported fields
}
```

File 代表一个打开的文件描述符。

### func Create

```go

```

### func CreateTemp

```go

```


### func NewFile

```go

```


### func Open

```go

```


### func OpenFile

```go

```


### func (*File) Chdir

```go

```


### func (*File) Chmod

```go

```


### func (*File) Chown

```go

```


### func (*File) Close

```go

```


### func (*File) Fd

```go

```


### func (*File) Name

```go

```


### func (*File) Read

```go

```

### func (*File) ReadAt

```go

```

### func (*File) ReadDir

```go

```

### func (*File) ReadFrom

```go

```

### func (*File) Readdir

```go

```

### func (*File) Readdirnames

```go

```

### func (*File) Seek

```go

```

### func (*File) SetDeadline

```go

```

### func (*File) SetReadDeadline

```go

```

### func (*File) SetWriteDeadline

```go

```

### func (*File) Stat

```go

```

### func (*File) Sync

```go

```

### func (*File) SyscallConn

```go

```

### func (*File) Truncate

```go

```

### func (*File) Write

```go

```

### func (*File) WriteAt

```go

```

### func (*File) WriteString

```go

```

### type FileInfo

```go

```

### func Lstat

```go

```

### func Stat

```go

```

### type FileMode

```go

```

### type LinkError

```go

```

### func (*LinkError) Error

```go

```

### func (*LinkError) Unwrap

```go

```

### type PathError

```go

```

### type ProcAttr

```go

```

### type Process

```go

```

### func FindProcess

```go

```

### func StartProcess

```go

```

### func (*Process) Kill

```go

```

### func (*Process) Release

```go

```

### func (*Process) Signal

```go

```

### func (*Process) Wait

```go

```

### type ProcessState

```go

```

### func (*ProcessState) ExitCode

```go

```

### func (*ProcessState) Exited

```go

```

### func (*ProcessState) Pid

```go

```

### func (*ProcessState) String

```go

```

### func (*ProcessState) Success

```go

```

### func (*ProcessState) Sys

```go

```

### func (*ProcessState) SysUsage

```go

```

### func (*ProcessState) SystemTime

```go

```

### func (*ProcessState) UserTime

```go

```

### type Signal

```go

```

### type SyscallError

```go

```

### func (*SyscallError) Error

```go

```

### func (*SyscallError) Timeout

```go

```

### func (*SyscallError) Unwrap

```go

```

## 目录

### [exec](os/exec.md)
### [signal](os/signal.md)
### [user](os/user.md)
