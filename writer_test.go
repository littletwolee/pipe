package pipe

import (
	"fmt"
	"testing"
)

func Test_Weffective(t *testing.T) {
	capacity := 4
	rBarrier := 0
	wBarrier := 0
	fmt.Println((rBarrier < 0 && wBarrier < capacity-_DEFAULT_ONE) || (0 <= rBarrier && rBarrier < wBarrier && wBarrier < capacity) || (wBarrier < rBarrier && rBarrier < capacity))
}
