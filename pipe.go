package pipe

import (
	"sync"
	"time"
)

type Pipe struct {
	pip             *pip
	stop, cleanStop chan bool
	jobs            *jobs
	wg              *sync.WaitGroup
}

func NewPipe(chNum int) *Pipe {
	return &Pipe{
		pip:       newPip(chNum),
		stop:      make(chan bool),
		cleanStop: make(chan bool),
		jobs: &jobs{
			m: new(sync.RWMutex),
		},
		wg: new(sync.WaitGroup),
	}
}

func (p *Pipe) AddJobs(jobs ...Job) {
	p.wg.Add(len(jobs))
	p.jobs.push(jobs...)
}

func (p *Pipe) Len() int {
	return p.jobs.len()
}

func (p *Pipe) Clean() {
	p.jobs.clean()
}
func (p *Pipe) cleanCache() {
	for {
		select {
		case <-p.cleanStop:
			break
		default:
			p.jobs.cleanCache()
			time.Sleep(1 * time.Minute)
		}
	}
}
func (p *Pipe) Wait() {
	time.Sleep(3 * time.Second)
	p.wg.Wait()
}

func (p *Pipe) Start(objs ...interface{}) {
	var obj interface{}
	if len(objs) > 0 {
		obj = objs[0]
	}
	go p.cleanCache()
	go func(p *Pipe, obj interface{}) {
		for {
			select {
			case j := <-p.pip.jobCH:
				p.pip.pipCH <- true
				go func(j Job, p *Pipe) {
					j.CallBack(j.Do(obj))
					p.wg.Done()
					<-p.pip.pipCH
				}(j, p)
			case <-p.pip.stopCH:
				p.pip.close()
				return
			}
		}
	}(p, obj)
	for {
		select {
		case <-p.stop:
			return
		default:
			j := p.jobs.pop()
			if j == nil {
				continue
			}
			p.pip.jobCH <- j
		}
	}
}

func (p *Pipe) Close() {
	p.jobs.clean()
	go func() { p.cleanStop <- true }()
	p.stop <- true
	p.pip.stopCH <- true
	close(p.stop)
}

type pip struct {
	jobCH  chan Job
	pipCH  chan bool
	stopCH chan bool
}

func newPip(chNum int) *pip {
	return &pip{
		jobCH:  make(chan Job, chNum),
		pipCH:  make(chan bool, chNum),
		stopCH: make(chan bool),
	}
}

func (p *pip) close() {
	close(p.jobCH)
	close(p.pipCH)
	close(p.stopCH)
}
