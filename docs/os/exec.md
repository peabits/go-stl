
# exec

## 概述

exec 包运行外部命令。它包装了 os.StartProcess 以便更轻松地重新映射标准输入和标准输出、使用管道连接 I/O 以及进行其他调整。

与来自 C 和其他语言的 “system” 库调用不同，os/exec 包有意不调用系统 shell 并且不扩展任何 glob 模式或处理通常由 shell 完成的其他扩展、管道或重定向。该包的行为更像是 C 的“exec”函数族。要扩展 glob 模式，要么直接调用 shell，注意转义任何危险的输入，要么使用 path/filepath 包的 Glob 函数。要扩展环境变量，请使用 package os 的 ExpandEnv。

请注意，此包中的示例假定为 Unix 系统。它们可能无法在 Windows 上运行，也无法在 golang.org 和 godoc.org 使用的 Go Playground 中运行。

### 当前目录下的可执行文件

函数 Command 和 LookPath 在当前路径中列出的目录中查找程序，遵循主机操作系统的约定。几十年来，操作系统一直将当前目录包含在此搜索中，有时默认情况下是隐式的，有时是显式配置的。现代的做法是，包含当前目录通常是意想不到的，而且常常会导致安全问题。

为了避免这些安全问题，从 Go 1.19 开始，这个包将不会解析使用相对于当前目录的隐式或显式路径条目的程序。也就是说，如果您运行 exec.LookPath("go")，无论路径如何配置，它都不会在 Unix 上成功返回 ./go，在 Windows 上也不会成功返回 .\go.exe。相反，如果通常的路径算法会产生该答案，则这些函数返回一个错误 err 满足 errors.Is(err, ErrDot)。

例如，考虑以下两个程序片段：

```go
path, err := exec.LookPath("prog")
if err != nil {
    log.Fatal(err)
}
use(path)
```

```go
cmd := exec.Command("prog")
if err := cmd.Run(); err != nil {
    log.Fatal(err)
}
```

无论当前路径如何配置，它们都不会找到并运行 ./prog 或 .\prog.exe。

总是想从当前目录运行程序的代码可以重写为 “./prog” 而不是 “prog”。

坚持包含相对路径条目结果的代码可以使用 errors.Is 检查来覆盖错误：

```go
path, err := exec.LookPath("prog")
if errors.Is(err, exec.ErrDot) {
    err = nil
}
if err != nil {
    log.Fatal(err)
}
use(path)
```

```go
cmd := exec.Command("prog")
if errors.Is(cmd.Err, exec.ErrDot) {
    cmd.Err = nil
}
if err := cmd.Run(); err != nil {
    log.Fatal(err)
}
```

设置环境变量 GODEBUG=execerrdot=0 会完全禁用 ErrDot 的生成，暂时恢复无法应用更有针对性的修复程序的 pre-Go 1.19 行为。 Go 的未来版本可能会删除对该变量的支持。

在添加此类覆盖之前，请确保您了解这样做的安全隐患。有关详细信息，请参阅 https://go.dev/blog/path-security。

## 索引

- [Constants](#常量)
- [Variables](#变量)
- [Functions](#函数)
- [Types](#类型)

## 常量

## 变量

## 函数

## 类型

## 目录

[fdtest](internal/fdtest.md)