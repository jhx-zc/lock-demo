package demo2

import "sync"

type Locker struct {
	l1 sync.Mutex
	//真实的锁
	l2 sync.Mutex

	locked bool
}

//
//  tryLock
//  @return bool true:got lock success  false:got lock failed
//
func (m *Locker) tryLock() bool {
	m.l1.Lock()
	defer m.l1.Unlock()
	if m.locked {
		return false
	}

	m.locked = true
	m.l2.Lock()
	return true
}

func (m *Locker) releaseLock() {
	m.l2.Unlock()
	m.l1.Lock()
	m.locked = false
	m.l1.Unlock()
}
