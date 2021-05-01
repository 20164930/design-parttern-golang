package state
// 对有状态的对象，把复杂的“判断逻辑”提取到不同的状态对象中，允许状态对象在其内部状态发生改变时改变其行为。

//状态模式是一种对象行为型模式，其主要优点如下。
//1.结构清晰，状态模式将与特定状态相关的行为局部化到一个状态中，并且将不同状态的行为分割开来，满足“单一职责原则”。
//2.将状态转换显示化，减少对象间的相互依赖。将不同的状态引入独立的对象中会使得状态转换变得更加明确，且减少对象间的相互依赖。
//3.状态类职责明确，有利于程序的扩展。通过定义新的子类很容易地增加新的状态和转换。
//
//状态模式的主要缺点如下。
//1.状态模式的使用必然会增加系统的类与对象的个数。
//2.状态模式的结构与实现都较为复杂，如果使用不当会导致程序结构和代码的混乱。
//3.状态模式对开闭原则的支持并不太好，对于可以切换状态的状态模式，增加新的状态类需要修改那些负责状态转换的源码，
//否则无法切换到新增状态，而且修改某个状态类的行为也需要修改对应类的源码。

type State interface {
	Handle() string
	getScore() int
}

type concreteState1 struct {
	score int
}
type concreteState2 struct {
	score int
}
type initialState struct {
	score int
}

type context struct{
	state State
}

func (c *concreteState1) Handle() string {
	return "state is state1, handle by state1"
}
func (c *concreteState1) getScore() int {
	return c.score
}
func (c *concreteState2) Handle() string {
	return "state is state2, handle by state2"
}
func (c *concreteState2) getScore() int {
	return c.score
}

func (c *initialState) Handle() string {
	return "state is state1, handle by state1"
}
func (c *initialState) getScore() int {
	return c.score
}

func (c *context) setState(s State) {
	c.state = s
}

func (c *context) Handle() string {
	return c.state.Handle()
}

func (c *context) addScore(num int) {
	score := c.state.getScore()
	score += num
	// 50分以下为状态1，以上为状态2
	switch {
	case score <=50:
		c.setState(&concreteState1{score})
	case score > 50:
		c.setState(&concreteState2{score})
	}
}