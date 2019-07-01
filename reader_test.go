package pipe

import (
	"fmt"
	"testing"
)

func Test_Reffective(t *testing.T) {
	capacity := 4
	rBarrier := 1
	wBarrier := 3
	fmt.Println((rBarrier < wBarrier && wBarrier < capacity) || (wBarrier < rBarrier && rBarrier < capacity))
}
