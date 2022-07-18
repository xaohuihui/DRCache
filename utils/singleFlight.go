package utils

import "sync"

// author: xaohuihui
// datetime: 2022/2/15 15:51:17
// software: GoLand
// singleFight.go 防止多次转发请求  降低从节点收到请求频率

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do 无论 Do 被调用多少次， 函数fn 都只会调用一次， 等待fn调用结束，返回返回值或错误
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	// 若该key的请求存在
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() // 若请求正在进行中，则等待，阻塞， 直到锁被释放
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1) // 发起请求前 锁加1
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done() // 锁减1

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
