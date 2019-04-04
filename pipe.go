package pipe

import (
	"sync"
)

type Pipe struct {
	pip  *pip
	stop chan bool
	jobs *jobs
}

func NewPipe(chNum int) *Pipe {
	return &Pipe{
		pip:  newPip(chNum),
		stop: make(chan bool),
		jobs: &jobs{
			m: new(sync.RWMutex),
		},
	}
}

func (p *Pipe) AddJobs(jobs ...Job) {
	p.jobs.push(jobs...)
}

func (p *Pipe) Len() int {
	return p.jobs.len()
}

func (p *Pipe) Clean() {
	p.jobs.clean()
}

func (p *Pipe) Start(objs ...interface{}) {
	var obj interface{}
	if len(objs) > 0 {
		obj = objs[0]
	}
	go func(p *pip, obj interface{}) {
		for {
			select {
			case j := <-p.jobCH:
				p.pipCH <- true
				go func(j Job, p *pip) {
					j.CallBack(j.Do(obj))
					<-p.pipCH
				}(j, p)
			case <-p.stopCH:
				p.close()
				return
			default:
			}
		}
	}(p.pip, obj)
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
