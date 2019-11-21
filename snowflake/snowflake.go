package snowflake

import (
	"errors"
	"sync"
	"time"
)

// 原始代码来源于 https://github.com/holdno/snowFlakeByGo
// 因为snowFlake目的是解决分布式下生成唯一id 所以ID中是包含集群和节点编号在内的

// const (
// 	workerBits uint8 = 10 // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点
// 	numberBits uint8 = 12 // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位数 即每毫秒可生成 2^12-1=4096个唯一ID
// 	// 这里求最大值使用了位运算，-1 的二进制表示为 1 的补码，感兴趣的同学可以自己算算试试 -1 ^ (-1 << nodeBits) 这里是不是等于 1023
// 	workerMax   int64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
// 	numberMax   int64 = -1 ^ (-1 << numberBits) // 同上，用来表示生成id序号的最大值
// 	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
// 	workerShift uint8 = numberBits              // 节点ID向左的偏移量
// 	// 41位字节作为时间戳数值的话 大约68年就会用完
// 	// 假如你2010年1月1日开始开发系统 如果不减去2010年1月1日的时间戳 那么白白浪费40年的时间戳啊！
// 	// 这个一旦定义且开始生成ID后千万不要改了 不然可能会生成相同的ID
// 	epoch int64 = 1525705533000 // 这个是我在写epoch这个变量时的时间戳(毫秒)
// )

type Options struct {
	Epoch       int64 // 初始时间戳，毫秒
	WorkerBits  uint8 // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点
	NumberBits  uint8 // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位数 即每毫秒可生成 2^12-1=4096个唯一ID
	workerMax   int64 // 节点ID的最大值，用于防止溢出
	numberMax   int64 // 同上，用来表示生成id序号的最大值
	timeShift   uint8 // 时间戳向左的偏移量
	workerShift uint8 // 节点ID向左的偏移量
}

// 默认worker参数
func DefaultOptions() Options {
	return NewOptions(1574326989581, 10, 12)
}

// 创建worker参数
func NewOptions(epoch int64, workerBits uint8, numberBits uint8) Options {
	return Options{
		Epoch:       epoch,
		WorkerBits:  workerBits,
		NumberBits:  numberBits,
		workerMax:   -1 ^ (-1 << workerBits),
		numberMax:   -1 ^ (-1 << numberBits),
		timeShift:   workerBits + numberBits,
		workerShift: numberBits,
	}
}

// 定义一个woker工作节点所需要的基本参数
type Worker struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录时间戳
	workerId  int64      // 该节点的ID
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
	options   Options    // 参数
}

var InvalidWorkerID = errors.New("worker ID excess of quantity")

// 实例化一个工作节点
func NewWorker(workerId int64, options Options) (*Worker, error) {
	// 要先检测workerId是否在上面定义的范围内
	if workerId < 0 || workerId > options.workerMax {
		return nil, InvalidWorkerID
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
		options:   options,
	}, nil
}

// 接下来我们开始生成id
// 生成方法一定要挂载在某个woker下，这样逻辑会比较清晰 指定某个节点生成id
func (w *Worker) GetId() int64 {
	// 获取id最关键的一点 加锁 加锁 加锁
	w.mu.Lock()
	defer w.mu.Unlock() // 生成完成后记得 解锁 解锁 解锁

	// 获取生成时的时间戳
	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if w.timestamp == now {
		w.number++

		// 这里要判断，当前工作节点是否在1毫秒内已经生成numberMax个ID
		if w.number > w.options.numberMax {
			// 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.number = 0
		w.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	}

	// 第一段 now - epoch 为该算法目前已经奔跑了xxx毫秒
	// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	id := (now-w.options.Epoch)<<w.options.timeShift | (w.workerId << w.options.workerShift) | (w.number)
	return id
}
