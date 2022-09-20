## go_mutex

[http://legendtkl.com/2016/10/23/golang-mutex/](http://legendtkl.com/2016/10/23/golang-mutex/)

[https://boilingfrog.github.io/2021/03/14/sync.Mutex/](https://boilingfrog.github.io/2021/03/14/sync.Mutex/)

[Understanding Real-World Concurrency Bugs in Go](https://songlh.github.io/paper/go-study.pdf)

### 同步原理的适用场景

1. 共享资源。并发地读写共享资源，会出现数据竞争（data race）的问题，所以需要 Mutex、RWMutex 这样的并发原语来保护。
2. 任务编排。需要 goroutine 按照一定的规律执行，而 goroutine 之间有相互等待或者依赖的顺序关系，我们常常使用 WaitGroup 或者 Channel 来实现。
3. 消息传递。信息交流以及不同的 goroutine 之间的线程安全的数据交流，常常使用 Channel 来实现。

互斥锁(Mutex)是使用最广泛的同步原语 Synchronization primitives，是信号量(Semaphore)的一种

### 互斥锁的实现机制

**临界区**: 在并发编程中，如果程序中的一部分会被并发访问或修改，
那么，为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，
这部分被保护起来的程序，就叫做临界区。

使用互斥锁，限定临界区只能同时由一个线程持有

### 问题

如果 Mutex 已经被一个 goroutine 获取了锁，其它等待中的 goroutine 们只能一直等待。那么，等这个锁释放后，等待中的 goroutine 中哪一个会优先获取 Mutex 呢？

等待的goroutine们是以FIFO排队的
1）当Mutex处于正常模式时，若此时没有新goroutine与队头goroutine竞争，则队头goroutine获得。若有新goroutine竞争大概率新goroutine获得。
2）当队头goroutine竞争锁失败1ms后，它会将Mutex调整为饥饿模式。进入饥饿模式后，锁的所有权会直接从解锁goroutine移交给队头goroutine，此时新来的goroutine直接放入队尾。
3）当一个goroutine获取锁后，如果发现自己满足下列条件中的任何一个#1它是队列中最后一个#2它等待锁的时间少于1ms，则将锁切换回正常模式
以上简略翻译自https://golang.org/src/sync/mutex.go 中注释Mutex fairness.

### mutex实现

**CAS**: CAS 指令将给定的值和一个内存地址中的值进行比较，如果它们是同一个值，就使用新值替换内存地址中的值，这个操作是原子性的。
原子性保证这个指令总是基于最新的值进行计算，如果同时有其它线程已经修改了这个值，那么，CAS 会返回失败。

**Unlock 方法可以被任意的 goroutine 调用释放锁，即使是没持有这个互斥锁的 goroutine，也可以进行这个操作。
这是因为，Mutex 本身并没有包含持有这把锁的 goroutine 的信息，所以，Unlock 也不会对此进行检查。
Mutex 的这个设计一直保持至今。**

### 位运算
```
&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空（AND NOT）
<<     左移
>>     右移
```
**&**
参与运算的两数各对应的二进位相与，两个二进制位都为1时，结果才为1
```
    0101
AND 0011
  = 0001
```

**|**
参与运算的两数各对应的二进位相或，两个二进制位都为1时，结果才为0
```
   0101（十进制5）
OR 0011（十进制3）
 = 0111（十进制7）
```

**^**
按位异或运算，对等长二进制模式或二进制数的每一位执行逻辑异或操作。操作的结果是如果某位不同则该位为1，否则该位为0
```
    0101
XOR 0011
  = 0110
```

**&^**
将运算符左边数据相异的位保留，相同位清零
```
   0001 0100 
&^ 0000 1111 
 = 0001 0000  
```

**<<**
各二进位全部左移若干位，高位丢弃，低位补0
```
   0001（十进制1）
<< 3（左移3位）
 = 1000（十进制8）
```

**>>**
各二进位全部右移若干位，对无符号数，高位补0，有符号数，各编译器处理方法不一样，有的补符号位（算术右移），有的补0
```
   1010（十进制10）
>> 2（右移2位）
 = 0010（十进制2）
```

### 错误用法
1. Lock/Unlock 不是成对出现
2. Copy 已使用的 Mutex，Mutex 是一个有状态的对象，它的 state 字段记录这个锁的状态。
3. 重入，Mutex 不是可重入的锁，因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。理论上，任何 goroutine 都可以随意地 Unlock 这把锁
















