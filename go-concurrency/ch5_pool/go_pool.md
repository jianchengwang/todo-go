## go_pool

Go 是一个自动垃圾回收的编程语言，采用[三色并发标记算法](https://go.dev/blog/ismmkeynote)标记对象并回收

Go 的自动垃圾回收机制还是有一个 STW（stop-the-world，程序暂停）的时间，而且，大量地创建在堆上的对象，也会影响垃圾回收标记的时间

### sync.pool

1. sync.Pool，它池化的对象可能会被垃圾回收掉，这对于数据库长连接等场景是不合适的
2. sync.Pool 本身就是线程安全的，多个 goroutine 可以并发地调用它的方法存取对象；
3. sync.Pool 不可在使用之后再复制使用。
4. 注意内存泄露跟内存浪费

### 第三方库

1. [bytebufferpool](https://github.com/valyala/bytebufferpool)
2. [oxtoacart/bpool](https://github.com/oxtoacart/bpool)
3. [gammazero/workerpool](https://pkg.go.dev/github.com/gammazero/workerpool)
4. [ivpusic/grpool](https://pkg.go.dev/github.com/ivpusic/grpool)
5. [dpaks/goworkers](https://pkg.go.dev/github.com/dpaks/goworkers)
