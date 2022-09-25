# go-number

go-number用于在API中兼容字符串和数值类型的数字。

当api的上下游采用不同语言编写时，
可能存在同一个数值字段有时候带引号有时候不带引号。

此外由于go语言的类型是带有初始值的，
不太容易判断json原文中是`0`还是空。

因此设计了`Number`类型，
示例见`example/main.go`，
规则见`number/number_test.go`。

```go
type Number struct {
Str string
Val int
}
```

## 行为

- 在序列化时忽略`Str`字段，只根据`Val`将`Number`序列化为数组类型。
- 在反序列化时，首先将json中的数据去掉双引号(如果有的话)保存到`Str`字段， 然后尝试转换为整数保存到`Val`字段。

## 限制

- 目前只支持整数
- 字段类型不能是指针`*Number`
