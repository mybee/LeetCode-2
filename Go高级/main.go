package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan int, 1)
	go func() {
		<-a
		fmt.Println("lllll")
	}()
	close(a)
	time.Sleep(10 * time.Second)

}

// OSI七层协议
// 应用层 -> 传输层 -> 网络层 -> 链路层

// 应用层 表示层, 会话层
// 传输层
// 网络层
// 链路层 物理层

// 临界 资源

//一个groutine睡觉的时候, 会被别的groutine 抢占资源
//

// 互斥锁:

//在使用互斥锁的时候, 使用完一定要解锁

// 读写锁

//	多个读操作可以同时进行
//  上了写锁之后, 不能读也不能写

// 非缓冲通道 不用写容量

// 缓冲通道 make 写上容量

// 闭包
//

// goroutine 池

//

// 单向通道
//  一般用在函数的参数, 对外只写 或 只读的 通道

// CSP 并发模型

//  调度机制 GPM 模型

// goroutine 是用户态的线程, 其他的是操作系统线程,
// 系统线程最小2m, 有固定的栈内存
// 协程2k, 栈不是固定的, 可以动态增大缩小, 最大可达一个G
//

//  G 就是goroutine, 存一些goroutine信息, 还有和p的绑定信息

//  P 管理一组goroutine队列, P里面会存储
//    当前goroutine运行的上下文, P会对自己管理的
//    goroutine队列做一些调度,  比如把(占用CPU时间较长的go协程暂停, 运行后续的go协程
//    当自己队列的消费完了, 会去全局队列里去取, 如果全
//    局队列的goroutine也消费完了, 会去其他的P的队列里抢任务.

//  M 是 go运行时对操作系统内核线程的虚拟, M 和 内核线程一般是一一映射的关系 物理线程
//  go协程最终是要放到M上执行的

//  P 和 M 也是一一对应的, 当一个G长久阻塞在M上时, runtime会新建一个M,
//  阻塞G所在的P会把其他的G挂载到新建的M上, 当阻塞的G完成或死亡, 回收旧的M

//  P的个数通过 runtime.GOMAXPROCS()去设定, 设置太多会频繁切换
//  m:n的调度技术

// goroutine 用户态的线程
// 优点:
//   协程调度是在用户态完成的, 不涉及系统内核态和用户态的频繁切换, 调用成本
//  很低,
//   利用了多核的资源

// 原子操作
// 	多协程计数使用原子++

// sync.Once
// 	声明 .do 在多协程只执行一次

// context
//	with parent
//	with deadline
//	with cancel

// etcd
// 1. 配置中心
//		配置热加载
// 2. 分布式锁
//		类似redis的setnx
// 奇数集群部署
//  etcd 图示: http://thesecretlivesofdata.com/raft/

// 性能优化
//  CPU
// 	Memory
// 	死锁 阻塞
// 	goroutine使用情况

// pprof
// 将cpu使用情况写到.pprof文件中
// 然后 go tool pprof xx.pprof
// top 10 查看cpu使用最多的10个函数
// 参数:
// 		flat: 当前函数占用cpu的耗时
// 		flat%: 当前函数占用cpu耗时百分比
// 		sum%: 耗时累计百分比
//		cum: 当前函数加上调用函数占用时间
// list + 函数名 可以查看代码耗时情况

// Raft 算法
//
// 自动选主流程
// 1. 根据超时时间, 最短的变候选者, 然后发起投票, 投票多的会变成leader
// 2. 如果leader节点挂了, 会重新选举, 前leader活了之后, 自动变成跟随者

// 一致性算法
//
// 数据同步:
//    leader先写需要更新日志, 将日志发给别的节点, 然后同时更新数据
// 	  如果网络分区, 又合并, leader 的小弟多的 数据为主

// redis 连接池
// 	 最大空闲数:
//   最大连接数: 0表示没有限制
//

// 布隆过滤器
//
