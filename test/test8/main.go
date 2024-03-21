package main

/*
分布式 ID 生成器通常可以使用 Snowflake 算法
通过将 ID 分解为不同的部分，包括时间戳、机器 ID、数据中心 ID 和序列号等部分
*/
import (
	"fmt"
	"sync"
	"time"
)

// SnowflakeIDGenerator 实现 Snowflake 算法的 ID 生成器
type SnowflakeIDGenerator struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	sequence      int64
}

// NewSnowflakeIDGenerator 创建一个新的 SnowflakeIDGenerator 实例
func NewSnowflakeIDGenerator(workerID int64) *SnowflakeIDGenerator {
	return &SnowflakeIDGenerator{
		workerID: workerID,
	}
}

// Generate 生成一个新的分布式 ID
func (g *SnowflakeIDGenerator) Generate() int64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	timestamp := time.Now().UnixNano() / 1e6 // 毫秒级时间戳

	if timestamp == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & 4095 // 12 位序列号，和4096-1作与运算取其低 12 位的值
		if g.sequence == 0 {
			// 序列号溢出，等待下一个毫秒
			for timestamp <= g.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		g.sequence = 0
	}

	g.lastTimestamp = timestamp

	// 构造 Snowflake ID
	id := ((timestamp - Epoch) << 22) | (g.workerID << 12) | g.sequence
	return id
}

// Epoch 是 Snowflake 算法的起始时间戳
const Epoch = 1609459200000 // 2021-01-01 00:00:00

func main() {
	// 创建一个 workerID 为 1 的 ID 生成器
	generator := NewSnowflakeIDGenerator(1)

	// 生成 5 个分布式 ID
	for i := 0; i < 5; i++ {
		id := generator.Generate()
		fmt.Println("Generated ID:", id)
	}
}
