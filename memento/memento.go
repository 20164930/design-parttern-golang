package memento
//备忘录（Memento）模式的定义：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，
//以便以后当需要时能将该对象恢复到原先保存的状态。该模式又叫快照模式。

//备忘录模式是一种对象行为型模式，其主要优点如下。
//1.提供了一种可以恢复状态的机制。当用户需要时能够比较方便地将数据恢复到某个历史的状态。
//2.实现了内部状态的封装。除了创建它的发起人之外，其他对象都不能够访问这些状态信息。
//3.简化了发起人类。发起人不需要管理和保存其内部状态的各个备份，所有状态信息都保存在备忘录中，并由管理者进行管理，这符合单一职责原则。
//
//其主要缺点是：资源消耗大。如果要保存的内部状态信息过多或者特别频繁，将会占用比较大的内存资源

//发起人
type Originator struct {
	state string
}

type Memento struct {
	state string
}
// 管理员
type Caretaker struct {
	memento Memento
}

func(o *Originator) setState(state string) {
	o.state = state
}
func(o *Originator) getState() string {
	return o.state
}
func(o *Originator) createMemento() Memento {
	return Memento{o.state}
}

func(o *Originator) restoreMemento(m Memento){
	o.state = m.state
}
func(m *Memento) setState(state string) {
	m.state = state
}

func(m *Memento) getState() string {
	return m.state
}
func(c *Caretaker) setMemento(m Memento) {
	c.memento = m
}
func(c *Caretaker) getMemento() Memento {
	return c.memento
}
