package pipe

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_WorkShop(t *testing.T) {
	var wg sync.WaitGroup
	ti := time.Now().UnixNano()
	pro := 1000000
	//pro := 10
	num := 100
	chNum := 2
	pip, err := NewPipe(int64(chNum), int64(num))
	if err != nil {
		t.Error(err)
	}

	//para := &para{sex: "male"}
	pip.Start()
	for index := 0; index < pro; index++ {
		wg.Add(1)

		pip.Write(newTestJob(index, &wg))
		//fmt.Printf("index:%d\n", index)
		//fmt.Println(pip.ringBuffer)
	}
	fmt.Println(111)
	// for {
	// 	// c := pip.cursors[0]
	// 	// next := c.r.getNext()
	// 	// fmt.Printf("w:%d,r:%d,eff:%v,n:%d\n", c.w.barrier.load(), c.r.barrier.load(), c.r.effective(next, c.w), next)
	// 	// //fmt.Println(c.r.defaultBarrier)
	// 	// time.Sleep(time.Second)
	// 	// if pip.Len() == 0 {
	// 	// 	pip.Stop()
	// 	// 	break
	// 	// }
	// }

	wg.Wait()
	ti = (time.Now().UnixNano() - ti) / 1000000
	if ti == 0 {
		ti = 1
	}
	fmt.Printf("opsPerSecond: %d\n", int64(pro*1000)/ti)

}

type para struct {
	sex string
}

func (t *testJob) NotNil() bool {
	return t != nil
}

type testJob struct {
	id int
	wg *sync.WaitGroup
}

func newTestJob(id int, wg *sync.WaitGroup) Job {
	t := new(testJob)
	t.id = id
	t.wg = wg
	if t == nil {
		fmt.Println("111111111111")
	}
	return t
}

func (t *testJob) Do() error {
	//fmt.Printf("a:%d\n", t.id)
	t.wg.Done()
	return nil
}
func (t *testJob) CallBack(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
