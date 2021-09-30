package demo1

import (
	"lock_demo/jobs"
	"sync"
)

var l = &Locker{}

func DoWork(wg *sync.WaitGroup, threadId int) {
	jobs.Job1(threadId)
	jobs.Job2(threadId)
	jobs.Job3(threadId)
	jobs.Job4(threadId)

	l.doLock()
	jobs.Job5(threadId)
	l.doUnLock()
	wg.Done()
}
