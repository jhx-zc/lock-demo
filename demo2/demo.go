package demo2

import (
	"go.uber.org/atomic"
	"lock_demo/jobs"
	"sync"
	"time"
)

var l1 = &Locker{locked: atomic.NewBool(false)}
var l2 = &Locker{locked: atomic.NewBool(false)}
var l3 = &Locker{locked: atomic.NewBool(false)}
var l4 = &Locker{locked: atomic.NewBool(false)}
var l5 = &Locker{locked: atomic.NewBool(false)}

type Jb struct {
	j func(threadId int)

	l *Locker
}

type worker struct {
	threadId int
	jobList  map[int]*Jb
}

func (m *worker) doJob() {
	//delete(m.jobList, 1)
	for len(m.jobList) > 0 {
		isFree := true
		for k, v := range m.jobList {
			if v.l.tryLock() {
				v.j(m.threadId)
				v.l.releaseLock()
				delete(m.jobList, k)
				isFree = false
				break
			}
		}
		if isFree {
			time.Sleep(1 * time.Second)
		}
	}
}

func NewWorker(threadId int) *worker {
	jl := make(map[int]*Jb)
	jl[1] = &Jb{
		j: jobs.Job1,
		l: l1,
	}
	jl[2] = &Jb{
		j: jobs.Job2,
		l: l2,
	}
	//jl[3] = &Jb{
	//	j: jobs.Job3,
	//	l: l3,
	//}
	//jl[4] = &Jb{
	//	j: jobs.Job4,
	//	l: l4,
	//}
	//jl[5] = &Jb{
	//	j: jobs.Job5,
	//	l: l5,
	//}
	//不用争抢锁资源
	jl[6] = &Jb{
		j: jobs.Job3,
		l: &Locker{locked: atomic.NewBool(false)},
	}
	jl[7] = &Jb{
		j: jobs.Job4,
		l: &Locker{locked: atomic.NewBool(false)},
	}
	jl[8] = &Jb{
		j: jobs.Job5,
		l: &Locker{locked: atomic.NewBool(false)},
	}

	return &worker{jobList: jl, threadId: threadId}
}

func DoWork(wg *sync.WaitGroup, threadId int) {
	w := NewWorker(threadId)
	w.doJob()
	if wg != nil {
		wg.Done()
	}
}
