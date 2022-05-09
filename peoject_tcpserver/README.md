已实现功能:公聊、私聊、重命名、超时强踢

---

运行服务端

window: go build -o server.exe main.go server.go client.go

linux: go build -o server main.go server.go client.go

运行客户端

window: go build -o client.exe client.go

linux: go build -o client client.go

---

#### defer

函数什么周期结束时, 执行 defer 所指定的语句.

```go
defer foo(params)
```

一个函数可以有多个 defer , 先 defer 的后调用.(栈)

defer 在 return 之后执行.

#### go

类似线程,但比线程效率更高

* 创建协程

```go
go foo(params)
```

* 创建匿名协程

```go
go func(params){
    statements
}(params)
```

* 退出协程

```go
runtime.Goexit()
```

如果协程中用了死循环,可用 `runtime.Goexit()` 或 `return` 退出该协程

#### chan `channel`

用于 go 协程之间传递消息

* 创建通道

```go
channel := make(chan datatype,capacity)
```

datatype:数据类型;capacity:容量 (留空默认为1).

* 长度 & 容量 获取

```go
len(channel)
cap(channel)
```

* 写入 & 读出

```go
channel <- value // 写入
value <-channel // 读出
```

往 channel 写入值时,channel 的 len 增加,在 len = capacity 时写入,会发生阻塞,直到另1个协程从 channel 中读取值时,才继续运行,才进行写入.

往 channel 读取值时,channel 的 len 减少,在 len = 0 时读取,会发生阻塞,直到另1个协程往 channel 写入值时,才继续运行,进行读取.

注:读出语句中,`<-channel` 中没有空格.

#### select

类似 switch 的语法,但和 switch 完全不同.

select 的作用监控多个 channel 的状态,哪个 channel 不阻塞,就执行哪个 case.

```go
select{
case channel1 <- value:
    break
case val2<-channel2:
    break
case <-channel3:
    break
}
```

