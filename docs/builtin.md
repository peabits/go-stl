
# builtin

buildin 包为 Go 的预声明标识符提供了文档。这里记录的项目实际上并不在 buildin 包中，但它们在这里的描述允许 godoc 为语言的特殊标识符提供文档。

## Constants

true 和 false 是两个无类型的布尔值。

```go
const (
    true = 0 == 0
    false = 0 != 0
)
```

iota 是一个预先声明的标识符，表示（通常用括号括起来的）const 声明中当前 const 规范的无类型整数序号。它是零索引的。

```go
const iota = 0
```

## Variables

nil 是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型的零值。

```go
var nil Type
```

## Functions

func append

```go
func append(slice []Type, elems ...Type) []Type
```

func cap

```go
func cap(v Type) int
```

func close

```go
func close(c chan <- Type>)
```

func complex

```go
func complex(r, i FloatType) ComplexType
```

func copy

```go
func copy(dst, src []Type) int
```

func delete

```go
func delete(m map[Type]Type1, key Type)
```

func imag

```go
func imag(c ComplexType) FloatType
```

func len

```go
func len(v Type) int
```

func make

```go
func make(t Type, size ...IntegerType) Type
```

func new

```go
func new(Type) *Type
```

func panic

```go
func panic(v any)
```

func print

```go
func print(args ...Type)
```

func println

```go
func println(args ...Type)
```

func real

```go
func read(c ComplexType) FloatType
```

func recover

```go
func recover() any
```

## Types

