package observer

import "fmt"

//意图：定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
//主要解决：一个对象状态改变给其他对象通知的问题，而且要考虑到易用和低耦合，保证高度的协作。

//其主要优点如下。
//1.降低了目标与观察者之间的耦合关系，两者之间是抽象耦合关系。符合依赖倒置原则。
//2. 目标与观察者之间建立了一套触发机制。
//
//它的主要缺点如下。
//1. 目标与观察者之间的依赖关系并没有完全解除，而且有可能出现循环引用。
//2. 当观察者对象很多时，通知的发布会花费很多时间，影响程序的效率。

type Observer interface{
	getName() string
	response()
}

type Subject interface {
	getObserverList() []Observer
	add(Observer)
	remove(Observer)
	notifyAll()
}

type ConcreteObserver1 struct {
	name string
}
type ConcreteObserver2 struct {
	name string
}

type ConcreteSubject struct{
	ObserverList []Observer
}

func (c *ConcreteObserver1) getName() string{
	return c.name
}
func (c *ConcreteObserver2) getName() string{
	return c.name
}


func (c *ConcreteObserver1) response() {
	fmt.Printf("observer1: %s known\n", c.name)
}
func (c *ConcreteObserver2) response() {
	fmt.Printf("observer2: %s known\n", c.name)
}

func (c *ConcreteSubject) getObserverList() []Observer{
	return c.ObserverList
}
func (c *ConcreteSubject) add(o Observer) {
	c.ObserverList = append(c.ObserverList, o)
}

func (c *ConcreteSubject) remove(o Observer) {
	for index := range c.ObserverList {
		if c.ObserverList[index].getName() == o.getName() {
			c.ObserverList = append(c.ObserverList[:index], c.ObserverList[index+1:]...)
			break
		}
	}
}
func (c *ConcreteSubject) notifyAll() {
	for _,v := range c.ObserverList {
		v.response()
	}
}
