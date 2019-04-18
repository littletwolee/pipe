package pipe

import (
	"fmt"
	"sync"
)

var max = 1000

type jobs struct {
	m    *sync.RWMutex
	n    int
	list list
}

type list []Job

func (l list) len() int {
	return len(l)
}

type Job interface {
	Do(obj interface{}) error
	CallBack(err error)
}

func (js *jobs) cleanCache() {
	js.m.Lock()
	defer js.m.Unlock()
	defer func(js *jobs) {
		if r := recover(); r != nil {
			fmt.Printf("error: %s, n: %d, len: %d\n", r, js.n, len(js.list))
		}
	}(js)
	if js.list.len() >= js.n {
		js.list = js.list[js.n:]
	}
	js.n = 0
}

func (js *jobs) pop() Job {
	js.m.Lock()
	defer js.m.Unlock()
	if js.list.len() > js.n {
		js.n++
		return js.list[js.n-1]
	}
	return nil
}

func (js *jobs) len() int {
	js.m.RLock()
	defer js.m.RUnlock()
	return js.list.len() - js.n
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
