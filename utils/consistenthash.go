package utils

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// author: xaohuihui
// datetime: 2022/1/20 14:29:45
// software: GoLand

/*  一致性哈希算法，
	1、计算虚拟节点hash值，放在环上
	2、计算key的hash值，再环上顺时针寻找到应选取的虚拟节点
{真实节点 ：[虚拟节点...]...}
*/

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constants all hashed keys
type Map struct {
	hash     Hash
	replicas int            // 多少个虚拟节点
	keys     []int          // Sorted
	hashMap  map[int]string // 虚拟节点与真实节点的映射关系
}

// New 新建一个 Map 实例，fn可以是自定义hash函数，不传默认为 crc32.ChecksumIEEE算法
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 新增一些keys去哈希,允许传入0或多个真实节点名称， 对每个真实节点key，创建m.replicas个虚拟节点，
// 虚拟节点名称为 strconv.Itoa(i) + key, 通过添加编号的方式区分不同虚拟节点
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			// 虚拟节点的哈希值，添加到环上
			m.keys = append(m.keys, hash)
			// 增加虚拟节点和真实节点的映射关系
			m.hashMap[hash] = key
		}
	}
	// 对环上的hash值排序
	sort.Ints(m.keys)
}

// Get 选择节点方法
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// 二分查找 适合的虚拟节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 若idx == len(m.keys) 说明应选择 m.keys[0]， 所以用求余的方式选择坐标
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
