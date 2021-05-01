package strategy
// 在软件开发中也常常遇到类似的情况，当实现某一个功能存在多种算法或者策略，
//我们可以根据环境或者条件的不同选择不同的算法或者策略来完成该功能，
//如数据排序策略有冒泡排序、选择排序、插入排序、二叉树排序等

//模式定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的变化不会影响使用算法的客户
//策略模式的主要优点如下。
//1.多重条件语句不易维护，而使用策略模式可以避免使用多重条件语句，如 if...else 语句、switch...case 语句。
//2.策略模式提供了一系列的可供重用的算法族，恰当使用继承可以把算法族的公共代码转移到父类里面，从而避免重复的代码。
//3.策略模式可以提供相同行为的不同实现，客户可以根据不同时间或空间要求选择不同的。
//4.策略模式提供了对开闭原则的完美支持，可以在不修改原代码的情况下，灵活增加新算法。
//5.策略模式把算法的使用放到环境类中，而算法的实现移到具体策略类中，实现了二者的分离。
//
//其主要缺点如下。
//1.客户端必须理解所有策略算法的区别，以便适时选择恰当的算法类。
//2.策略模式造成很多的策略类，增加维护难度。

type strategy interface {
	execMethod() string
}
type context struct {
	strategy
}
type concreteStrategy1 struct{}
type concreteStrategy2 struct{}

func (c *concreteStrategy1) execMethod() string {
	return "strategy1 has been exec"
}
func (c *concreteStrategy2) execMethod() string {
	return "strategy2 has been exec"
}
func (c *context) setStrategy(s strategy) {
	c.strategy = s
}
func (c *context) execStrategy() string {
	return c.strategy.execMethod()
}