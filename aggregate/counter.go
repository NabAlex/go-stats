package aggregate

import "sync"

type Counter struct {
    sync.Mutex
    tag   string
    count int
}

func (c *Counter) GetVal() int {
    return c.count
}

func (c *Counter) GetName() string {
    return c.tag
}

func (c *Counter) Up() {
    c.Mutex.Lock()
    defer c.Mutex.Unlock()
    c.count++
}

func (c *Counter) Refresh() {
    c.Mutex.Lock()
    defer c.Mutex.Unlock()
    c.count = 0
}

func CreateCounter(tag string) Aggregator {
    c := new(Counter)
    c.tag = tag
    c.count = 0
    return c
}
