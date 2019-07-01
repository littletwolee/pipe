package pipe

// type writers struct {
// 	w        *writer
// 	capacity int64
// }

func newWriter(rchunk, chunk int64, capacity int64) *writer {
	return &writer{barrier: newBarrier(rchunk), chunk: chunk, defaultBarrier: _DEFAULT_ZERO, capacity: capacity}
}

type writer struct {
	barrier                         *barrier
	defaultBarrier, capacity, chunk int64
}

func (w *writer) effective(wBarrier int64, r *reader) bool {
	rBarrier := r.barrier.load()
	return (rBarrier < wBarrier && wBarrier < r.capacity) || (wBarrier < rBarrier && rBarrier < r.capacity)
}

func (w *writer) next(current int64) {
	next := current + w.chunk
	if next < w.capacity {
		w.barrier.store(next)
		return
	}
	w.barrier.store(w.defaultBarrier)
}
