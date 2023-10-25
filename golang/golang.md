# golang

## 基础

### slice

1. slice 的底层结构

```golang
type slice struct {
    array unsafe.Pointer // 指向底层数组的指针
    len int // 当前元素个数
    cap int // 当期的容量
}
```

2. slice 的扩容
   // src/runtime/slice.go/growslice

- 期望最小容量时候大于原容量的 2 倍
  - 是，使用期望容量
  - 否，根据版本判断
    - go 1.9 判断 slice 的原长度是否小于 1024，是，则新容量等于 2 倍的原容量，否，新容量等于 1.25 倍原容量
    - go 1.16 判断 slice 的原长度是否小于 1024，是，则新容量等于 2 倍的原容量，否，新容量等于 1.25 倍原容量，执导容量大于期望容量
    - go 1.18 判断 slice 的原长度是否小于 256，是，则新容量等于 2 倍的原容量，否，newcap = oldcap+（oldcap+ 3\*256）/4 原容量，执导容量大于期望容量

3. slice 和 array 的区别

- slice 可以扩容，array 数量固定
-

### map

1. map 的底层结构

```golang
type Map struct {
	key, elem Type
}

```

2. map 是并发安全的吗？怎样并发安全？
   map 不是并发安全的。参考[map.go](./map.go)当多个线程同时写时会提示**fatal error: concurrent map writes**
   map 并发操作的方案：
   - 读写锁
   - sync.Map
3. map 的扩容

4. map 读取是顺序的吗？怎样实现顺序读取？
   map 的读取不是顺序的，如果要实现顺序读取，需要使用 slice 缓存 map 的 key，然后通过遍历 slice 的方式获得有序的 key，参考[有序 map](./map_seq.go)

5. sync.Map

### channel

1. channel 的底层结构

2. 生产者和消费者
   参考[生产者和消费者](./producer_consumer.go)

## go 的调度和并发

1. GMP 模型
   G： goroutine 轻量线程

2. 进程，线程和协程

进程：一个程序就是一个进程
线程：
协程：

3. goroutine 飙升如何排查

4.

### 内存逃逸以及避免措施

struct 是否可以比较？什么类型不能比较

- 不同类型的 struct 之间不能进行比较，编译期就会报错
- 同类型的 struct 也分为两种情况，
  - struct 的所有成员都是可以比较的，则该 strcut 的不同实例可以比较
  - struct 中含有不可比较的成员（如 Slice），则该 struct 不可以比较

不可比较的类型：slice，map

线上程序 CPU / 内存 占用过高定位，排查与解决
(https://www.infoq.cn/article/f69uvzjuomq276hbp1qb)
(https://xie.infoq.cn/article/a94a46596a14cc6578d7e8ad4)
cpu 占用 99% -> 发现 GC 线程占用率持续异常 -> 怀疑是内存问题 -> 排查对象数量 -> 定位产生对象异常多的接口 -> 定位到某接口 -> 在日志中找到此接口的异常请求 -> 根据异常参数排查代码中的问题 -> 定位到问题
程序性能问题分析

runtime/pprof (用于工具类型，一次性)
net/http/pprof （用于服务类型，一直运行）
