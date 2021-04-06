package command
//命令（Command）模式的定义如下：将一个请求封装为一个对象，使发出请求的责任和执行请求的责任分割开。
//这样两者之间通过命令对象进行沟通，这样方便将命令对象进行储存、传递、调用、增加与管理。
//何时使用：在某些场合，比如要对行为进行"记录、撤销/重做、事务"等处理，这种无法抵御变化的紧耦合是不合适的。
//优点：
//1、降低了系统耦合度。
//2、新的命令可以很容易添加到系统中去。
//
//缺点：
//1.使用命令模式可能会导致某些系统有过多的具体命令类。

type Invoker struct{
	c command
}

type command interface {
	execute() string
}
type concreteCommandA struct {
	executor executorA
}

type concreteCommandB struct {
	executor executorB
}
type executorA struct {}
type executorB struct {}

func(c *concreteCommandA) execute() string{
	return c.executor.doAction()
}
func(c *concreteCommandB) execute() string{
	return c.executor.doAction()
}
func(e *executorA) doAction() string{
	return "executorA doAction success"
}
func(e *executorB) doAction() string{
	return "executorB doAction success"
}
func (i *Invoker) setCommand(commandStr string) {
	switch commandStr {
	case "A":
		i.c=&concreteCommandA{}
	case "B":
		i.c=&concreteCommandB{}
	}
}
func (i *Invoker) call() string{
	return i.c.execute()
}
func createInvoker() Invoker {
	return Invoker{}
}