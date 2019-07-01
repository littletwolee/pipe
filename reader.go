package pipe

type readers []*reader

func newReaders(chunk int64, capacity int64) readers {
	rs := make(readers, chunk)
	for i := int64(_DEFAULT_ZERO); i < chunk; i++ {
		rs[i] = newReader(i, chunk, capacity)
	}
	return rs
}
func newReader(chunkNum, chunk int64, capacity int64) *reader {
	return &reader{barrier: newBarrier(chunkNum), chunk: chunk, defaultBarrier: chunkNum, capacity: capacity}
}

type reader struct {
	barrier                         *barrier
	chunk, defaultBarrier, capacity int64
}

func (r *reader) effective(rBarrier int64, w *writer) bool {
	wBarrier := w.barrier.load()
	return (rBarrier <= wBarrier && wBarrier < r.capacity) || (wBarrier < rBarrier && rBarrier < r.capacity)
}

func (r *reader) next(current int64) {
	next := current + r.chunk
	if next < r.capacity {
		r.barrier.store(next)
		return
	}
	r.barrier.store(r.defaultBarrier)
}
