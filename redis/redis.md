# redis

## 基础

1. redis 的数据结构
   redis 常用的数据结构包括 String， List， Hash， Set，ZSet（Sorted Set）
   除此之外，还包括 GEO，HyperLogLog，bitmap 等

2. String 的常用命令及用途

```shell
set xxx
get xxx
setex
setnx
```

用途：缓存数据，分布式锁

3. List 的常用命令及用途

```shell
lpush
rpop
rpush
lpop
lterm
```

用途：消息队列

4. Hash 的常用命令及用途

```shell
hset [key] [field] [value]
hget [key] [field]
hmset [key] [field1] [field2]
```

用途：缓存数据

5. Set 的常用命令及用途

```shell

```

6. ZSet 的常用命令及用途

```shell

```

## 缓存雪崩，缓存击穿和缓存穿透

缓存雪崩

缓存击穿

缓存穿透

## 分布式锁

## redis 线程模型

redis 是单线程吗？如果请求量过大会怎样？
redis 从接收客户端请求，解析请求，执行查询/写入 这个过程是单线程的，但是 redis 本身不是单线程的，还有 BIO 后台线程，用于处理文件关闭，flush 等操作

## redis 的复制
