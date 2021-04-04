package proxy

//代理模式的定义：由于某些原因需要给某对象提供一个代理以控制对该对象的访问。
//这时，访问对象不适合或者不能直接引用目标对象，代理对象作为访问对象和目标对象之间的中介。
//在代码中，一般代理会被理解为代码增强，实际上就是在原代码逻辑前后增加一些代码逻辑，而使调用者无感知。

//代理模式的主要优点有：
//1.代理模式在客户端与目标对象之间起到一个中介作用和保护目标对象的作用；
//2.代理对象可以扩展目标对象的功能；
//3.代理模式能将客户端与目标对象分离，在一定程度上降低了系统的耦合度，增加了程序的可扩展性
//
//其主要缺点是：
//1.代理模式会造成系统设计中类的数量增加
//2.在客户端和目标对象之间增加一个代理对象，会造成请求处理速度变慢；
//3.增加了系统的复杂度；

//抽象主题
type subject interface {
	solveRequest() string
}
//真实主题
type RealSubject struct {}
func(r *RealSubject) solveRequest() string {
	return "solve request complete"
}

type proxy struct {
	subject *RealSubject
}
//代理在处理前后做一些额外的处理
func (p *proxy) solveRequest() string {
	pre := p.doPreRequest()
	after := p.doAfterRequest()
	return pre + p.subject.solveRequest() + after
}

func (p *proxy) doPreRequest()string {
	return "["
}
func (p *proxy) doAfterRequest()string {
	return "]"
}

func createProxy() subject {
	return &proxy{&RealSubject{}}
}