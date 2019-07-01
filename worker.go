package pipe

var (
	_chNum int
)

type workers struct {
	ch chan *worker
}

func newWorkers(chNum int, bufferMax int64) *workers {
	w := make(chan *worker, chNum)
	for i := -1; i < chNum-1; i++ {
		w <- &worker{defaultChunk: int64(i), chunk: int64(i)}
	}
	_chNum = chNum
	return &workers{
		ch: w,
	}
}

type worker struct {
	defaultChunk, chunk int64
}

// func (w *worker) load(js *jobs) Job {
// 	return js.load(w, w.swap(w.chunk+int64(_chNum), js.bufferMax))
// }

func (w *worker) swap(i, bufferMax int64) int64 {
	if i > bufferMax {
		w.chunk = w.defaultChunk
		return w.defaultChunk
	}
	return i
}
