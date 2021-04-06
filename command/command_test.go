package command

import "testing"

func TestCommand(t *testing.T) {
	invoker := createInvoker()
	invoker.setCommand("A")
	ret := invoker.call()
	if ret != "executorA doAction success" {
		t.Error("commandA execute failed")
	}
	invoker.setCommand("B")
	ret = invoker.call()
	if ret != "executorB doAction success" {
		t.Error("commandB execute failed")
	}
}
