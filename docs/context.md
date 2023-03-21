
# context

包上下文定义了上下文类型，它携带截止日期、取消信号和其他跨 API 边界和进程之间的请求范围的值。

对服务器的传入请求应该创建一个上下文，对服务器的传出调用应该接受一个上下文。它们之间的函数调用链必须传播上下文，可选择将其替换为使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 创建的派生上下文。当一个上下文被取消时，从它派生的所有上下文也被取消。

WithCancel、WithDeadline 和 WithTimeout 函数采用 Context（父级）并返回派生的 Context（子级）和 CancelFunc。调用 CancelFunc 会取消子项及其子项，删除父项对子项的引用，并停止任何关联的计时器。未能调用 CancelFunc 会泄漏子级及其子级，直到父级被取消或计时器触发。 go vet 工具检查是否在所有控制流路径上使用了 CancelFuncs。

WithCancelCause 函数返回一个 CancelCauseFunc，它接受错误并将其记录为取消原因。在取消的上下文或其任何子上下文上调用 Cause 来检索原因。如果没有指定原因，Cause(ctx) 返回与 ctx.Err() 相同的值。

使用 Contexts 的程序应该遵循这些规则，以保持接口在包之间的一致性，并启用静态分析工具来检查上下文传播：

不要将上下文存储在结构类型中；相反，将 Context 显式传递给需要它的每个函数。 Context 应该是第一个参数，通常命名为 ctx：
