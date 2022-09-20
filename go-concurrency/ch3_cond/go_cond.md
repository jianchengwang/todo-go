## go_cond

等待 / 通知（wait/notify）机制

请实现一个限定容量的队列（queue），当队列满或者空的时候，利用等待 / 通知机制实现阻塞或者唤醒。

从开发实践上，我们真正使用 Cond 的场景比较少，因为一旦遇到需要使用 Cond 的场景，我们更多地会使用 Channel 的方式，因为那才是更地道的 Go 语言的写法，[issues/21165](https://github.com/golang/go/issues/21165)

### 基本方法

```go
type Cond
    func NeWCond(l Locker) *Cond
    func (c *Cond) Broadcast()
    func (c *Cond) Signal()
    func (c *Cond) Wait()
```

Cond 关联的 Locker 实例可以通过 c.L 访问，它内部维护着一个先入先出的等待队列。

[通用方法名](https://en.wikipedia.org/wiki/Monitor_(synchronization)#Condition_variables_2)

1. Signal 方法，允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，则需要从等待队列中移除第一个 goroutine 并把它唤醒。在其他编程语言中，比如 Java 语言中，Signal 方法也被叫做 notify 方法。调用 Signal 方法时，不强求你一定要持有 c.L 的锁。
2. Broadcast 方法，允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；如果 Cond 等待队列中有一个或者多个等待的 goroutine，则清空所有等待的 goroutine，并全部唤醒。在其他编程语言中，比如 Java 语言中，Broadcast 方法也被叫做 notifyAll 方法。同样地，调用 Broadcast 方法时，也不强求你一定持有 c.L 的锁。
3. Wait 方法，会把调用者 Caller 放入 Cond 的等待队列中并阻塞，直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒。

### 常见错误

1. Cond 最常见的使用错误，也就是调用 Wait 的时候没有加锁。
2. 只调用了一次 Wait

如果你想在使用 Cond 的时候避免犯错，只要时刻记住调用 cond.Wait 方法之前一定要加锁，以及 waiter goroutine 被唤醒不等于等待条件被满足这两个知识点。

