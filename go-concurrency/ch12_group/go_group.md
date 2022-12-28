## go_group

共享资源保护、任务编排和消息传递是 Go 并发编程中常见的场景，而分组执行一批相同的或类似的任务则是任务编排中一类情形

### ErrGroup

ErrGroup 就是用来应对这种场景的。它和 WaitGroup 有些类似，但是它提供功能更加丰富：
1. 和 Context 集成；
2. error 向上传播，可以把子任务的错误传递给 Wait 的调用者。

接下来，我来给你介绍一下 ErrGroup 的基本用法和几种应用场景。

```go

// 同时还会返回一个使用 context.WithCancel(ctx) 生成的新 Context。一旦有一个子任务返回错误，或者是 Wait 调用返回，这个新 Context 就会被 cancel。
func WithContext(ctx context.Context) (*Group, context.Context)


// 执行子任务的 Go 方法
func (g *Group) Go(f func() error)

// 类似 WaitGroup，Group 也有 Wait 方法，等所有的子任务都完成后，它才会返回,否则只会阻塞等待。如果有多个子任务返回错误，它只会返回第一个出现的错误，如果所有的子任务都执行成功，就返回 nil
func (g *Group) Wait() error

```
