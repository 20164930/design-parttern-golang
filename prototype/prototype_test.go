package prototype

import "testing"

func TestPrototype(t *testing.T) {
	manager := Manage{}
	productA := manager.create("A")
	productB := manager.create("B")

	copyA := productA.clone()
	copyB := productB.clone()
	if copyA.getName() != productA.getName() {
		t.Error("productA and copyA is different")
	}
	if copyB.getName() != productB.getName() {
		t.Error("productB and copyB is different")
	}
}
