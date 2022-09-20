## go_channel

### CSP模型

Communicating Sequential Process 的简称，中文直译为通信顺序进程

[CSP 最早出现于计算机科学家 Tony Hoare论文](https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf)

[go引入csp历史](https://swtch.com/~rsc/thread/)

CSP 允许使用进程组件来描述系统，它们独立运行，并且只通过消息传递的方式通信。

Channel 类型是 Go 语言内置的类型，你无需引入某个包，就能使用它

share memory by communicating”则是类似于 CSP 模型的方式，通过通信的方式，一个 goroutine 可以把数据的“所有权”交给另外一个 goroutine（虽然 Go 中没有“所有权”的概念，但是从逻辑上说，你可以把它理解为是所有权的转移）。

### 使用场景

1. 数据交流：当作并发的 buffer 或者 queue，解决生产者 - 消费者问题。多个 goroutine 可以并发当作生产者（Producer）和消费者（Consumer）
2. 数据传递：一个 goroutine 将数据交给另一个 goroutine，相当于把数据的拥有权 (引用) 托付出去。
3. 信号通知：一个 goroutine 可以将信号 (closing、closed、data ready 等) 传递给另一个或者另一组 goroutine 。
4. 任务编排：可以让一组 goroutine 按照一定的顺序并发或者串行的执行，这就是编排的功能。
5. 利用 Channel 也可以实现互斥锁的机制。

### 基本类型

这个箭头总是射向左边的，元素类型总在最右边。如果箭头指向 chan，就表示可以往 chan 中塞数据；如果箭头远离 chan，就表示 chan 会往外吐数据。

```go
chan string          // 可以发送接收string
chan<- struct{}      // 只能发送struct{}
<-chan int           // 只能从chan接收int
```

```go
chan<- （chan int） // <- 和第一个chan结合
chan<- （<-chan int） // 第一个<-和最左边的chan结合，第二个<-和左边第二个chan结合
<-chan （<-chan int） // 第一个<-和最左边的chan结合，第二个<-和左边第二个chan结合 
chan (<-chan int) // 因为括号的原因，<-和括号内第一个chan结合
```

### 初始化

通过 make，我们可以初始化一个 chan，未初始化的 chan 的零值是 nil。你可以设置它的容量，比如下面的 chan 的容量是 9527，我们把这样的 chan 叫做 buffered chan；
如果没有设置，它的容量是 0，我们把这样的 chan 叫做 unbuffered chan。

还有一个知识点需要你记住：nil 是 chan 的零值，是一种特殊的 chan，对值是 nil 的 chan 的发送接收调用者总是会阻塞。

```go
make(chan int, 9527)
```

### 发送数据

这里的ch是chan int类型 或者chan <-int

```go
ch <- 2000
```

### 接收数据

收数据时，还可以返回两个值。第一个值是返回的 chan 中的元素，很多人不太熟悉的是第二个值。
第二个值是 bool 类型，代表是否成功地从 chan 中读取到一个值，如果第二个参数是 false，chan 已经被 close 而且 chan 中没有缓存的数据，这个时候，第一个值是零值。
所以，如果从 chan 读取到一个零值，可能是 sender 真正发送的零值，也可能是 closed 的并且没有缓存元素产生的零值。

```go
  x := <-ch // 把接收的一条数据赋值给变量x
  foo(<-ch) // 把接收的一个的数据作为参数传给函数
  <-ch // 丢弃接收的一条数据
```

### 其他操作

Go 内建的函数 close、cap、len 都可以操作 chan 类型：close 会把 chan 关闭掉，cap 返回 chan 的容量，len 返回 chan 中缓存的还未被取走的元素数量。

send 和 recv 都可以作为 select 语句的 case clause，如下面的例子：

```go

func main() {
    var ch = make(chan int, 10)
    for i := 0; i < 10; i++ {
        select {
        case ch <- i:
        case v := <-ch:
            fmt.Println(v)
        }
    }
}
```

chan 还可以应用于 for-range 语句中，比如：

```go
for v := range ch {
    fmt.Println(v)
}
```

或者是忽略读取的值，只是清空 chan：

```go
for range ch {
}
```

### 常见的错误
1. close 为 nil 的 chan； -> panic
2. send 已经 close 的 chan；-> panic
3. close 已经 close 的 chan。-> panic
4. 内存泄露

### 建议
1. 共享资源的并发访问使用传统并发原语；
2. 复杂的任务编排和消息传递使用 Channel；
3. 消息通知机制使用 Channel，除非只想 signal 一个 goroutine，才使用 Cond；
4. 简单等待所有任务的完成用 WaitGroup，也有 Channel 的推崇者用 Channel，都可以；
5. 需要和 Select 语句结合，使用 Channel；
6. 需要和超时配合时，使用 Channel 和 Context。

### chan状态

![img.png](img.png)








