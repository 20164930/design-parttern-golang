package mediator

import "testing"

func TestMediator (t *testing.T) {
	mediator1 := ConcreteMediator{}
	colleague1 := ConcreteColleague1{id: "1"}
	colleague2 := ConcreteColleague1{id: "2"}
	mediator1.register(&colleague1)
	mediator1.register(&colleague2)

	colleague1.send("hello , I am 1")
	colleague2.send("hello , I am 2")
}
