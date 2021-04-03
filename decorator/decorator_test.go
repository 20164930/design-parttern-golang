package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	decorator1 := decoratorWithOperation2{&decorator{&concreteComponentA{}}}
	decorator2 := decoratorWithOperation3{&decorator{&concreteComponentA{}}}
	decorator3 := decoratorWithOperation2{&decorator{&concreteComponentB{}}}
	decorator4 := decoratorWithOperation3{&decorator{&concreteComponentB{}}}

	decorator1.operation()
	fmt.Println()
	decorator2.operation()
	fmt.Println()
	decorator3.operation()
	fmt.Println()
	decorator4.operation()
}
