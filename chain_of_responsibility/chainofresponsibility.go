package chain_of_responsibility

import "strings"

//责任链（Chain of Responsibility）模式的定义：
//为了避免请求发送者与多个请求处理者耦合在一起，于是将所有请求的处理者通过前一对象记住其下一个对象的引用而连成一条链；
//当有请求发生时，可将请求沿着这条链传递，直到有对象处理它为止。

//使用：在处理消息的时候做过滤器。 拦截的类都实现统一接口。
//优点：
//1、降低耦合度。它将请求的发送者和接收者解耦。
//2、简化了对象。使得对象不需要知道链的结构。
//3、增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任。
//4、增加新的请求处理类很方便。
//
//缺点：
//1、不能保证请求一定被接收。
//2、系统性能将受到一定影响，而且在进行代码调试时不太方便，可能会造成循环调用。
//3、可能不容易观察运行时的特征，有碍于除错。
type handler interface {
	setNext(handler)
	getNext() *handler
	solveRequest(string) string
}

type concreteHandler1 struct{
	next *handler
}
type concreteHandler2 struct{
	next *handler
}

func (c *concreteHandler1) setNext(h handler) {
	c.next = &h
}

func (c *concreteHandler1) getNext() *handler {
	return c.next
}

func (c *concreteHandler1) solveRequest(request string) string{
	if strings.EqualFold(request, "one") {
		return "response1"
	} else if c.next != nil {
		return (*c.next).solveRequest(request)
	}
	return "can not solve this request"
}

func (c *concreteHandler2) setNext(h handler) {
	c.next = &h
}

func (c *concreteHandler2) getNext() *handler {
	return c.next
}

func (c *concreteHandler2) solveRequest(request string) string{
	if strings.EqualFold(request, "two") {
		return "response2"
	} else if c.next != nil {
		return (*c.next).solveRequest(request)
	}
	return "can not solve this request"
}

func createChain() handler {
	handler1 := &concreteHandler1{}
	handler2 := &concreteHandler2{}
	handler1.setNext(handler2)
	return handler1
}