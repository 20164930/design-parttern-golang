package strategy

import "testing"

func TestStrategy(t *testing.T) {
	context := context{&concreteStrategy1{}}
	ret1 := context.execStrategy()
	if ret1 != "strategy1 has been exec" {
		t.Error("strategy1 error")
	}

	context.setStrategy(&concreteStrategy2{})
	ret2 := context.execStrategy()
	if ret2 != "strategy2 has been exec" {
		t.Error("strategy2 error")
	}
}
