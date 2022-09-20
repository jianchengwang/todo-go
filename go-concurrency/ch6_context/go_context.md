## go_context

### 使用场景
1. 上下文信息传递 （request-scoped），比如处理 http 请求、在请求处理链路上传递信息；
2. 控制子 goroutine 的运行；
3. 超时控制的方法调用；
4. 可以取消的方法调用。

### 生成方法

1. context.Background()：返回一个非 nil 的、空的 Context，没有任何值，不会被 cancel，不会超时，没有截止日期。一般用在主函数、初始化、测试以及创建根 Context 的时候。
2. context.TODO()：返回一个非 nil 的、空的 Context，没有任何值，不会被 cancel，不会超时，没有截止日期。当你不清楚是否该用 Context，或者目前还不知道要传递一些什么上下文信息的时候，就可以使用这个方法。

其实，你根本不用费脑子去考虑，可以直接使用 context.Background

### 约定俗成的规则

1. 一般函数使用 Context 的时候，会把这个参数放在第一个参数的位置
2. 从来不把 nil 当做 Context 类型的参数值，可以使用 context.Background() 创建一个空的上下文对象，也不要使用 nil
3. Context 只用来临时做函数之间的上下文透传，不能持久化 Context 或者把 Context 长久保存。把 Context 持久化到数据库、本地文件或者全局变量、缓存中都是错误的用法
4. key 的类型不应该是字符串类型或者其它内建类型，否则容易在包之间使用 Context 时候产生冲突。使用 WithValue 时，key 的类型应该是自己定义的类型
5. 常常使用 struct{}作为底层类型定义 key 的类型。对于 exported key 的静态类型，常常是接口或者指针。这样可以尽量减少内存分配。

### 创建特殊的Context
1. WithValue: 基于parent Context生成新的Context，保存一个key-value的键值对，常常用来传递上下文
2. WithCancel: 返回parent的副本，只是副本中的Done channel是新建的对象，他的类型是cancelCtx
3. WithTimeout: 其实和WithDeadline一样，只不过一个参数是超时时间，一个是截止时间
4. WithDeadline: 返回一个parent的副本，并且设置了一个不晚于参数d的截止时间，类型为timeCtx(或者cancelCtx)
