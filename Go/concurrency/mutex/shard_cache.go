package mutex

import (
	"hash/fnv"
	"sync"
)

type Shard struct {
	sync.RWMutex
	data map[string]interface{}
}

type ShardedCache struct {
	shards     []*Shard
	shardCount int
}

func NewShardedCache(shardCount int) *ShardedCache {
	shards := make([]*Shard, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = &Shard{data: make(map[string]interface{})}
	}
	return &ShardedCache{shards: shards, shardCount: shardCount}
}

func (c *ShardedCache) getShard(key string) *Shard {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return c.shards[uint(hash.Sum32())%uint(c.shardCount)]
}

func (c *ShardedCache) Set(key string, value interface{}) {
	shard := c.getShard(key)
	shard.Lock()
	shard.data[key] = value
	shard.Unlock()
}
