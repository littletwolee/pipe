package pipe

import (
	"reflect"
	"runtime"
)

type cursors map[int64]*cursor

func newCursors(wl, rl, capacity int64, rb *ringBuffer, s *bool) cursors {
	cs := make(cursors, rl)
	w := newWriter(rl, wl, capacity)
	for i := int64(_DEFAULT_ZERO); i < rl; i++ {
		cs[i] = newCursor(w, rl, i, capacity, rb, s)
	}
	return cs
}

func (cs cursors) start() {
	for _, c := range cs {
		go c.start()
	}
}

type cursor struct {
	w     *writer
	r     *reader
	rb    *ringBuffer
	swich *bool
}

func newCursor(w *writer, rl, ri, capacity int64, rb *ringBuffer, s *bool) *cursor {
	return &cursor{
		w:     w,
		r:     newReader(ri, rl, capacity),
		rb:    rb,
		swich: s,
	}
}
func (c *cursor) start() {
	i := -1
	for {
		i++
		if !*c.swich {
			break
		}
		current := c.r.barrier.load()
		if c.r.effective(current, c.w) {
			job := (*c.rb)[current]
			if job == nil {
				c.r.next(current)
				continue
			}
			if reflect.ValueOf(job).IsNil() {
				runtime.Gosched()
				continue
			}
			job.CallBack(job.Do())
			(*c.rb)[current] = nil
			c.r.next(current)
			return
		}
		// fmt.Printf("r:%d,w:%d\n", current, c.w.barrier.load())
		// if i > 4 {

		// 	panic(111)
		// }

	}
}

func (c *cursor) write(j Job) bool {
	current := c.w.barrier.load()
	if c.w.effective(current, c.r) {
		(*c.rb)[current] = j
		c.w.next(current)
		return true
	}
	return false
}
