package Adapter

import "testing"

func TestAdapter(t *testing.T) {
	adapteeA := &AdapteeA{"A"}
	adapteeB := &AdapteeB{"B"}
	retA := createTarget(adapteeA).use()
	retB := createTarget(adapteeB).use()

	if retA != "A" {
		t.Error("Adapter A use error")
	}
	if retB != "B" {
		t.Error("Adapter B use error")
	}
}