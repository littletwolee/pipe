package pipe

import (
	"fmt"
	"runtime"
)

//Pipe .
type Pipe struct {
	capacity, chunk int64
	cursors         cursors
	ringBuffer      *ringBuffer
	swich           *bool
	current         *barrier
}

//NewPipe create a new Pipe
func NewPipe(rn, capacity int64) (*Pipe, error) {
	if !checkEven(int64(rn)) && rn != 1 {
		return nil, fmt.Errorf(_ERROR_EVEN, "reader number")
	}
	if !checkEven(capacity) {
		return nil, fmt.Errorf(_ERROR_EVEN, "buffer max")
	}
	p := &Pipe{
		capacity:   capacity,
		current:    newBarrier(),
		chunk:      rn,
		ringBuffer: newBuffer(capacity),
		swich:      &_STOP,
	}
	p.cursors = newCursors(_DEFAULT_ONE, rn, capacity, p.ringBuffer, p.swich)
	return p, nil
}

//Start pipe
func (p *Pipe) Start() {
	*p.swich = _START
	p.cursors.start()
}

//Stop pipe
func (p *Pipe) Stop() {
	*p.swich = _STOP
}

//Write job
func (p *Pipe) Write(j Job) {
	i := -1
	for {
		i++
		// if !*p.swich {
		// 	break
		// }
		old := p.current.load()
		c := p.cursors[old%p.chunk]

		if c.write(j) {
			p.current.add()
			return
		}
		fmt.Printf("o:%d,r:%d,w:%d\n", old, c.r.barrier.load(), c.w.barrier.load())
		if i > 4 {
			panic(111)
		}

		runtime.Gosched()
	}
}

// func (p *Pipe) Push(jobs ...Job) {
// 	p.js.push(jobs...)
// }

// func (p *Pipe) Len() int64 {
// 	return p.js.len()
// }

// func (p *Pipe) Start(objs ...interface{}) {
// 	for {
// 		select {
// 		case w := <-p.workers.ch:
// 			job := w.load(p.js)
// 			if job == nil {
// 				go func(w *worker, p *Pipe) {
// 					runtime.Gosched()
// 					time.Sleep(time.Second)
// 					p.workers.ch <- w
// 				}(w, p)
// 				continue
// 			}
// 			job.CallBack(job.Do(objs))
// 			p.workers.ch <- w
// 		case <-p.stop:
// 			break
// 		}
// 	}
// }

// func (p *Pipe) Stop() {
// 	p.stop <- true
// }
