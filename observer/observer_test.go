package observer

import "testing"

func TestObserver(t *testing.T) {
	observer1 := ConcreteObserver1{"observer1"}
	observer2 := ConcreteObserver1{"observer2"}
	observer3 := ConcreteObserver2{"observer3"}

	subject := ConcreteSubject{}
	subject.add(&observer1)
	subject.add(&observer2)
	subject.add(&observer3)

	subject.notifyAll()
}
