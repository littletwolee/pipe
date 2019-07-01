package pipe

// func Test_len(t *testing.T) {
// 	buffer := int64(10)
// 	js := newJobs(buffer)
// 	js.wBarrier = 5
// 	js.rBarrier = 1
// 	l := js.len()
// 	if l != 3 {
// 		t.Errorf("len error r: %d,w: %d, len should be: %d", js.rBarrier, js.wBarrier, l)
// 	}
// 	js.rBarrier = 6
// 	js.wBarrier = 3
// 	l = js.len()
// 	if l != 6 {
// 		t.Errorf("len error r: %d,w: %d, len should be: %d", js.rBarrier, js.wBarrier, buffer-js.rBarrier+js.wBarrier)
// 	}
// }

// func Test_push(t *testing.T) {
// 	buffer := int64(10)
// 	js := newJobs(buffer)
// 	ts := &testStruct{}
// 	js.push(newTest(ts))
// 	if js.len() != 1 {
// 		t.Errorf("push error should be 1 nor %d", js.len())
// 	}
// }

// type testStruct struct{}

// func newTest(ts *testStruct) Job {
// 	return ts
// }
// func (t testStruct) Do(obj ...interface{}) error { return nil }

// func (t testStruct) CallBack(err error) {}

// func Test_effective(t *testing.T) {
// 	buffer := int64(10)
// 	js := newJobs(buffer)
// 	js.rBarrier = 0
// 	js.wBarrier = 1
// 	if !js.wEffective(js.wBarrier, js.rBarrier) {
// 		t.Errorf("effective 1 error should be true")
// 	}
// 	js.rBarrier = 0
// 	js.wBarrier = 9
// 	if !js.wEffective(js.wBarrier, js.rBarrier) {
// 		t.Errorf("effective 2 error should be true")
// 	}
// 	js.rBarrier = 0
// 	js.wBarrier = 0
// 	if js.wEffective(js.wBarrier, js.rBarrier) {
// 		t.Errorf("effective 3 error should be false")
// 	}
// 	js.rBarrier = 1
// 	js.wBarrier = 0
// 	if !js.wEffective(js.wBarrier, js.rBarrier) {
// 		t.Errorf("effective 4 error should be true")
// 	}
// }
