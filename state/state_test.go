package state

import "testing"

func TestState(t *testing.T) {
	context := context{&initialState{}}
	context.addScore(30)
	ret := context.Handle()
	if ret!="state is state1, handle by state1" {
		t.Error("handle state1 error")
	}
	context.addScore(30)
	ret = context.Handle()
	if ret!="state is state2, handle by state2" {
		t.Error("handle state2 error")
	}
}
