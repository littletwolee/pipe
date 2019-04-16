package pipe

import (
	"sync"
)

var max = 1000

type jobs struct {
	m    *sync.RWMutex
	n    int
	list []Job
}

type Job interface {
	Do(obj interface{}) error
	CallBack(err error)
}

func (js *jobs) cleanCache() {
	js.m.Lock()
	defer js.m.Unlock()
	js.list = js.list[js.n:]
	js.n = 0
}

func (js *jobs) pop() Job {
	js.m.Lock()
	defer js.m.Unlock()
	var job Job
	if len(js.list) > js.n {
		job = js.list[js.n]
		js.n++
	}
	return job
}

func (js *jobs) len() int {
	js.m.RLock()
	defer js.m.RUnlock()
	return len(js.list) - js.n
}

func (js *jobs) clean() {
	js.m.Lock()
	defer js.m.Unlock()
	js.list = []Job{}
}

func (js *jobs) push(jobs ...Job) {
	js.m.Lock()
	defer js.m.Unlock()
	js.list = append(js.list, jobs...)
}

type e int

const (
	_EOF e = iota
)

func (e e) Do(obj interface{}) error                                { return nil }
func (e e) CallBack(obj interface{}, f func(obj interface{}) error) {}
