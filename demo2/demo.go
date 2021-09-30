package demo2

import (
	"go.uber.org/atomic"
	"lock_demo/jobs"
	"sync"
)

var l = &Locker{locked: atomic.NewBool(false)}

func tryDoJob5(isJob5Done *atomic.Bool, threadId int) {
	//fmt.Println(threadId, " tryDoJob5")
	if !isJob5Done.Load() {
		if l.tryLock() {
			jobs.Job5(threadId)
			isJob5Done.Store(true)
			l.releaseLock()
		}
	}
}

func DoWork(wg *sync.WaitGroup, threadId int) {
	isJonDone := atomic.NewBool(false)
	tryDoJob5(isJonDone, threadId)
	jobs.Job1(threadId)

	tryDoJob5(isJonDone, threadId)
	jobs.Job2(threadId)

	tryDoJob5(isJonDone, threadId)
	jobs.Job3(threadId)

	tryDoJob5(isJonDone, threadId)
	jobs.Job4(threadId)

	if !isJonDone.Load() {
		l.l2.Lock()
		jobs.Job5(threadId)
		l.l2.Unlock()
	}
	if wg != nil {
		wg.Done()
	}
}
