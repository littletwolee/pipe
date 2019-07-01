package pipe

import (
	"sync/atomic"
)

type barrier struct {
	b *int64
}

func newBarrier(bs ...int64) *barrier {
	if len(bs) <= 0 {
		temp := int64(_DEFAULT_ZERO)
		return &barrier{&temp}
	}
	return &barrier{&bs[0]}
}

func (b *barrier) load() int64 {
	return atomic.LoadInt64(b.b)
}

func (b *barrier) add(i ...int64) {
	if len(i) <= 0 {
		atomic.AddInt64(b.b, _DEFAULT_ONE)
		return
	}
	atomic.AddInt64(b.b, i[0])
}

func (b *barrier) compareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(b.b, old, new)
}

func (b *barrier) store(new int64) {
	atomic.StoreInt64(b.b, new)
}

// func effective(c *cursor) bool {
// 	abs := math.Abs(float64(c.r.getNext() - c.w.barrier.load()))
// 	return 0 < abs && abs < float64(c.r.capacity)
// }
