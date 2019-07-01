package pipe

const (
	_DEFAULT_MAX_BUFFER = 1024 * 100
)

// type disruptor struct {
// 	ringBuffer         []interface{}
// 	wBarrier, rBarrier barrier
// }
// type barrier int64

// func (b *barrier) int() int64 {
// 	return int64(*b)
// }

// func (b *barrier) load() int64 {
// 	return atomic.LoadInt64(b)
// }

// func (b *barrier) change(c int64) {
// 	atomic.AddInt64(b, c)
// }

// func New(max int64) *disruptor {
// 	return &disruptor{
// 		ringBuffer: make([]interface{}, max),
// 		rBarrier:   -1,
// 	}
// }

// func (d *disruptor) Push(o interface{}) {
// 	d.ringBuffer[d.wBarrier.load()] = o
// }
