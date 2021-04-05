package chain_of_responsibility

import "testing"

func TestChain(t *testing.T) {
	handlerIml := createChain()
	resp1 := handlerIml.solveRequest("one")
	resp2 := handlerIml.solveRequest("two")
	resp3 := handlerIml.solveRequest("three")
	if resp1 !="response1" {
		t.Error("solve request one error")
	}
	if resp2 !="response2" {
		t.Error("solve request two error")
	}
	if resp3 !="can not solve this request" {
		t.Error("solve request three error")
	}
}
