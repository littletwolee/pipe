package pipe

const (
	_DEFAULT_BUFFER_MAX = int64(1024 * 100)
)

type ringBuffer []Job

func newBuffer(max int64) *ringBuffer {
	rb := make(ringBuffer, pipeMax(max))
	return &rb
}

func pipeMax(pipeMax int64) int64 {
	if pipeMax > 0 && pipeMax < _DEFAULT_BUFFER_MAX {
		return pipeMax
	}
	return _DEFAULT_BUFFER_MAX
}

// type jobs struct {
// 	bufferMax, wBarrier, rBarrier int64
// 	list                          list
// }

// func newJobs(bufferMax int64) *jobs {
// 	return &jobs{
// 		wBarrier:  1,
// 		bufferMax: bufferMax,
// 		list:      make([]Job, bufferMax),
// 	}
// }

// type list []Job

// func (j *jobs) len() int64 {
// 	w := atomic.LoadInt64(&j.wBarrier)
// 	r := atomic.LoadInt64(&j.rBarrier) + 1
// 	if !j.isCover() && r < w {
// 		return w - r
// 	} else if j.isCover() && r > w {
// 		return j.bufferMax - r + w
// 	}
// 	return 0
// }

type Job interface {
	Do() error
	CallBack(err error)
	//NotNil() bool
}

var _nil Job = (Job)(nil)

type nilJob struct{}

func (n *nilJob) Do() error          { return nil }
func (n *nilJob) CallBack(err error) {}

// func (js *jobs) push(jobs ...Job) {
// 	for _, job := range jobs {
// 		wCurrent := atomic.LoadInt64(&js.wBarrier)
// 		rCurrent := atomic.LoadInt64(&js.rBarrier)
// 		for {
// 			if js.wEffective(wCurrent, rCurrent) {
// 				js.list[wCurrent] = job
// 				js.swap(wCurrent+1, &js.wBarrier)
// 				break
// 			}
// 			runtime.Gosched()
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func (js *jobs) load(w *worker, i int64) Job {
// 	wb := atomic.LoadInt64(&js.wBarrier)
// 	for {
// 		if !js.rEffective(wb, i) {
// 			time.Sleep(time.Second)
// 			runtime.Gosched()
// 		}
// 		w.chunk = i
// 		js.swap(i, &js.rBarrier)
// 		break
// 	}
// 	return js.list[i]
// }

// func (js *jobs) swap(i int64, p *int64) {
// 	if i > js.bufferMax {
// 		atomic.SwapInt64(p, 0)
// 		return
// 	}
// 	atomic.SwapInt64(p, i)
// }

// func (js *jobs) isCover() bool {
// 	return atomic.LoadInt64(&js.wBarrier) < atomic.LoadInt64(&js.rBarrier)
// }
// func (js *jobs) rEffective(w, r int64) bool {
// 	return r <= w || (w < r && r < js.bufferMax)
// }
// func (js *jobs) wEffective(w, r int64) bool {
// 	return (r < w && w < js.bufferMax) || w < r
// }

type e int

const (
	_EOF e = iota
)

func (e e) Do(obj interface{}) error                                { return nil }
func (e e) CallBack(obj interface{}, f func(obj interface{}) error) {}
