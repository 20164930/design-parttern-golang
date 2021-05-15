package visitor

import (
	"testing"
)

func TestVisitor (t *testing.T) {
	list := []Element{
		&ConcreteElementA{},
		&ConcreteElementB{},
	}
	objectList := ObjectStruct{list}
	visitor1 := &ConcreteVisitorA{}
	visitor2 := &ConcreteVisitorB{}
	ret1 := objectList.accept(visitor1)
	if ret1 != "visitorA visit elementA operator success\nvisitorA visit elementB operator success\n" {
		t.Error("visitor1 failed")
	}
	ret2 := objectList.accept(visitor2)
	if ret2 != "visitorB visit elementA operator success\nvisitorB visit elementB operator success\n" {
		t.Error("visitor2 failed")
	}

}
