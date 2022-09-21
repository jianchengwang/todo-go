## go_singleflight

SingleFlight 是 Go 开发组提供的一个扩展并发原语。它的作用是，在处理多个 goroutine 同时调用同一个函数的时候，只让一个 goroutine 去调用这个函数，等到这个 goroutine 返回结果的时候，再把结果返回给这几个同时调用的 goroutine，这样可以减少并发调用的数量。

sync.Once 不是只在并发的时候保证只有一个 goroutine 执行函数 f，而是会保证永远只执行一次，而 SingleFlight 是每次调用都重新执行，并且在多个请求同时调用的时候只有一个执行。
它们两个面对的场景是不同的，sync.Once 主要是用在单次初始化场景中，而 SingleFlight 主要用在合并并发请求的场景中，尤其是缓存场景


