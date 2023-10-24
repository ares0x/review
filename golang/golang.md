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

### channel

1. channel 的底层结构

###

##

## 并发
