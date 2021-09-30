package demo2

import (
	"go.uber.org/atomic"
	"sync"
)

type Locker struct {
	l1 sync.Mutex
	//真实的锁
	l2 sync.Mutex

	locked *atomic.Bool
}

//
//  tryLock
//  @return bool true:got lock success  false:got lock failed
//
func (m *Locker) tryLock() bool {
	if m.locked.Load() {
		return false
	}

	m.l1.Lock()
	defer m.l1.Unlock()
	if m.locked.Load() {
		return false
	}

	m.locked.Store(true)
	m.l2.Lock()
	return true
}

func (m *Locker) releaseLock() {
	m.l2.Unlock()
	m.l1.Lock()
	m.locked.Store(false)
	m.l1.Unlock()
}
