package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	input := []interface{}{"a", "b", "c"}
	concreteAggregate := aggregateA{input}
	iterator := concreteAggregate.getIterator()
	for ; iterator.hasNext(); {
		fmt.Println(iterator.next())
	}
}
