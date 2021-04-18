package mediator

import "fmt"

//中介者模式（Mediator Pattern）用一个中介对象来封装一系列的对象交互，
//中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。

//优点：
//1、降低了类的复杂度，将一对多转化成了一对一。
//2、各个类之间的解耦。
//3、符合迪米特原则。
//
//缺点：
//1.中介者会庞大，变得复杂难以维护。

type Mediator interface {
	register(colleague Colleague)
	// 转发
	relay(msg string, sender Colleague)
}
type Colleague interface {
	getID() string
	setMedium(Mediator)
	receive(string)
	send(string)
}

type ConcreteMediator struct {
	colleagues []Colleague
}

func (m *ConcreteMediator) register(colleague Colleague) {
	if m.colleagues == nil {
		m.colleagues = make([]Colleague, 0)
		m.colleagues = append(m.colleagues, colleague)
		colleague.setMedium(m)
		return
	}
	for _,v := range m.colleagues {
		if v.getID() == colleague.getID() {
			colleague.setMedium(m)
			return
		}
	}
	m.colleagues = append(m.colleagues, colleague)
	colleague.setMedium(m)
}

func (m *ConcreteMediator) relay(msg string, sender Colleague) {
	for _, v := range m.colleagues {
		// 广播
		if v.getID() != sender.getID() {
			v.receive(msg)
		}
	}
}

type ConcreteColleague1 struct {
	id string
	mediator Mediator
}

func (c *ConcreteColleague1) getID() string {
	return c.id
}

func (c *ConcreteColleague1) setMedium(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague1) receive(s string) {
	fmt.Printf("%s receive message: %s\n", c.getID(), s)
}

func (c *ConcreteColleague1) send(s string) {
	fmt.Printf("%s send message: %s\n", c.getID(), s)
	// 广播
	c.mediator.relay(s, c)
}

type ConcreteColleague2 struct {
	id string
	mediator Mediator
}

func (c *ConcreteColleague2) getID() string {
	return c.id
}

func (c *ConcreteColleague2) setMedium(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague2) receive(s string) {
	fmt.Printf("%s receive message: %s\n", c.getID(), s)
}

func (c *ConcreteColleague2) send(s string) {
	fmt.Printf("%s send message: %s\n", c.getID(), s)
	// 广播
	c.mediator.relay(s, c)
}

