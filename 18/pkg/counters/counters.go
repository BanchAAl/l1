package counters

import "sync"

//SyncCounter конкурентный счётчик
type SyncCounter struct {
	i int
	m sync.Mutex
}

//Decr увеличить счётчик
func (c *SyncCounter) Decr() {
	c.m.Lock()
	c.i--
	c.m.Unlock()
}

//Incr увеличить счётчик
func (c *SyncCounter) Incr() {
	c.m.Lock()
	c.i++
	c.m.Unlock()
}

//Value получение значения счётчика
func (c *SyncCounter) Value() int {
	var rslt int
	c.m.Lock()
	rslt = c.i
	c.m.Unlock()

	return rslt
}
