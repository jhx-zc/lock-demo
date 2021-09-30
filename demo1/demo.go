package demo1

import (
	"lock_demo/jobs"
	"sync"
)

var l = &Locker{}

func DoWork(wg *sync.WaitGroup, threadId int) {
	l.doLock()
	jobs.Job1(threadId)
	l.doUnLock()

	l.doLock()
	jobs.Job2(threadId)
	l.doUnLock()

	//l.doLock()
	jobs.Job3(threadId)
	//l.doUnLock()

	//l.doLock()
	jobs.Job4(threadId)
	//l.doUnLock()

	//l.doLock()
	jobs.Job5(threadId)
	//l.doUnLock()

	if wg != nil {
		wg.Done()
	}
}
