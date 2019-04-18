package pipe

import (
	"fmt"
	"testing"
	"time"
)

func Test_WorkShop(t *testing.T) {
	num := 3
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	chNum := 1
	pip := NewPipe(chNum)
	para := &para{sex: "male"}
	go pip.Start(para)
	// fmt.Println("start")
	// time.Sleep(1 * time.Minute)
	// pip.Close()
	// fmt.Println("close")
	// time.Sleep(1 * time.Minute)
	for index := 0; index < num; index++ {
		//fmt.Println(index)
		pip.AddJobs(newTestJob(index))
		//time.Sleep(1 * time.Second)
	}
	pip.Wait()
	// pip.Close()
	// time.Sleep(10 * time.Second)
	// fmt.Println("a")
	// pip = NewPipe(chNum)
	// fmt.Println("b")
	// go pip.Start(para)
	// fmt.Println("c")
	// for index := 0; index < num; index++ {
	// 	pip.AddJobs(newTestJob(index))
	// 	//time.Sleep(1 * time.Second)
	// }
	// fmt.Printf("completed add %d\n", num/2)
	// time.Sleep(1 * time.Minute)
	// for index := num/2 - 1; index < num; index++ {
	// 	pip.AddJobs(newTestJob(index))
	// }
	time.Sleep(10 * time.Hour)
	pip.Close()
	fmt.Printf("Pipe close success\n")
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
