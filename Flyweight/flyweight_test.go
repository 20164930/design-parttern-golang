package Flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factory := flyweightFactory{nil}
	fw := factory.getFlyweight("A")
	ret := fw.operation()
	if ret != "flyweight-A,outFlyweight-A" {
		t.Error("flyweight error")
	}
}
