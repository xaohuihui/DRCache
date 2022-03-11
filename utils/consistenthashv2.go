package utils

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// author: xaohuihui
// datetime: 2022/3/10 16:21:55
// software: GoLand

type Consistent struct {
	// 排序的hash虚拟节点
	hashSortedNodes []uint32
	// 虚拟节点对应的真实节点信息
	circle map[uint32]string
	// 已绑定的节点
	nodes map[string]bool
	// map读写锁
	sync.RWMutex
	// 虚拟节点数
	virtualNodeCount int
}

func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) Add(node string, virtualNodeCount int) error {
	if node == "" {
		return nil
	}
	c.Lock()
	defer c.Unlock()

	if c.circle == nil {
		c.circle = map[uint32]string{}
	}
	if c.nodes == nil {
		c.nodes = map[string]bool{}
	}

	if _, ok := c.nodes[node]; ok {
		return errors.New("node already existed")
	}
	c.nodes[node] = true
	// 增加虚拟节点
	for i := 0; i < virtualNodeCount; i++ {
		virtualKey := c.hashKey(node + strconv.Itoa(i))
		c.circle[virtualKey] = node
		c.hashSortedNodes = append(c.hashSortedNodes, virtualKey)
	}

	// 虚拟节点排序
	sort.Slice(c.hashSortedNodes, func(i, j int) bool {
		return c.hashSortedNodes[i] < c.hashSortedNodes[j]
	})
	return nil
}

func (c *Consistent) GetNode(key string) string {
	c.RLock()
	defer c.RUnlock()

	hash := c.hashKey(key)
	i := c.getPosition(hash)

	return c.circle[c.hashSortedNodes[i]]
}

func (c *Consistent) getPosition(hash uint32) int {
	i := sort.Search(len(c.hashSortedNodes), func(i int) bool {
		return c.hashSortedNodes[i] >= hash
	})
	if i < len(c.hashSortedNodes) {
		if i == len(c.hashSortedNodes)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.hashSortedNodes) - 1
	}
}
