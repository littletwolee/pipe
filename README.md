# workshop

## Summary

Golang pipline with channal & waitgroup

## Usage

```
// in ur shell
go get github.com/littletwolee/pipe
```

```
// import it
import (
	"github.com/littletwolee/pipe"
)
// get new pipe, chNum is a num that max workers can working simultaneously
ws := NewPipe(chNum) 
// u need implementation job interface
// type job interface {
//      Do() error
//      CallBack(func() error)
// }
// like these:
type testJob struct {
    id int
}

func newTestJob(id int) job {
    return &testJob{
        id: id,
    }
}

func (t *testJob) Do() error {
    time.Sleep(1 * time.Second)
    if t.id%2 == 0 {
        fmt.Printf("mod 2==0: %d\n", t.id)
        return nil
    }

    return fmt.Errorf("error %d", t.id)
}
func (t *testJob) CallBack(f func() error) {
    if err := f(); err != nil {
        fmt.Println(err)
    }
}
// prepare your data
go ws.Start()
// add all jobs in pipe
for index := 0; index < num; index++ {
    ws.AddJobs(newTestJob(index))
}


....
ur next step code
....


// close it
pip.Close()
```
