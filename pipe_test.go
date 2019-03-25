package workshop

import (
	"fmt"
	"testing"
	"time"
)

func Test_WorkShop(t *testing.T) {
	num := 100
	chNum := 5
	pip := NewPipe(chNum)
	para := &para{sex: "male"}
	go pip.Start(para)
	for index := 0; index < num/2; index++ {
		pip.AddJobs(newTestJob(index))
		//time.Sleep(1 * time.Second)
	}
	fmt.Printf("completed add %d\n", num/2)
	time.Sleep(1 * time.Minute)
	for index := num/2 - 1; index < num; index++ {
		pip.AddJobs(newTestJob(index))
	}
	time.Sleep(10 * time.Second)
	pip.Close()
	fmt.Printf("pipe close success\n")
	for {
	}
}

type para struct {
	sex string
}

type testJob struct {
	id int
}

func newTestJob(id int) Job {
	return &testJob{
		id: id,
	}
}

func (t *testJob) Do(obj interface{}) error {
	time.Sleep(3 * time.Second)
	if t.id%2 == 0 {
		fmt.Printf("mod 2==0: %d, sex: %s\n", t.id, obj.(*para).sex)
		return nil
	}
	return fmt.Errorf("error %d", t.id)
}
func (t *testJob) CallBack(obj interface{}, f func(obj interface{}) error) {
	if err := f(obj); err != nil {
		fmt.Println(err)
	}
}
