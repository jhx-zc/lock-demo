package demo1

import "sync"

type Locker struct {
	l sync.Mutex
}

func (m *Locker) doLock() {
	m.l.Lock()
}

func (m *Locker) doUnLock() {
	m.l.Unlock()
}
