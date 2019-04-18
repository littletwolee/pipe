package pipe

import (
	"fmt"
	"testing"
	"time"
)

func Test_WorkShop(t *testing.T) {
	num := 3
	chNum := 1
	pip := NewPipe(chNum)
	go pip.Start()
	for index := 0; index < num; index++ {
		pip.AddJobs(newTestJob(index))
	}
	pip.Wait()
	pip.Close()
	time.Sleep(2 * time.Minute)

	pip = NewPipe(chNum)
	go pip.Start()
	for index := 0; index < num; index++ {
		pip.AddJobs(newTestJob(index))
	}
	pip.Wait()
	pip.Close()
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
	time.Sleep(1 * time.Second)
	// if t.id%2 == 0 {
	// 	fmt.Printf("mod 2==0: %d, sex: %s\n", t.id, obj.(*para).sex)
	// 	return nil
	// }
	// return fmt.Errorf("error %d", t.id)
	fmt.Printf("aaa:%d\n", t.id)
	return nil
}
func (t *testJob) CallBack(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
