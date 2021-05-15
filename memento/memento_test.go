package memento

import "testing"

func TestMemento(t *testing.T) {
	originator := Originator{"state1"}
	memento := originator.createMemento()
	originator.setState("state2")
	originator.restoreMemento(memento)
	if originator.getState() != "state1" {
		t.Error("restore originator state error")
	}
}
