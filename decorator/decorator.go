package decorator

import "fmt"

//主要解决：一般的，我们为了扩展一个类经常使用继承方式实现，由于继承为类引入静态特征，并且随着扩展功能的增多，子类会很膨胀。
//何时使用：在不想增加很多子类的情况下扩展类。
//如何解决：将具体功能职责划分，同时继承装饰者模式。
//主要优点有：
//1.装饰器是继承的有力补充，比继承灵活，在不改变原有对象的情况下，动态的给一个对象扩展功能，即插即用
//2.通过使用不用装饰类及这些装饰类的排列组合，可以实现不同效果
//3.装饰器模式完全遵守开闭原则
//
//其主要缺点是：
//1.装饰器模式会增加许多子类，过度使用会增加程序得复杂性。

//要扩展功能的对象的接口
type component interface {
	operation1()
}
//装饰器，负责给对象扩展功能，相当于在对象最外层封装了一层
type decorator struct {
	component
}

//待扩展功能接口的实现类
type concreteComponentA struct {}
type concreteComponentB struct {}

func(c *concreteComponentA) operation1() {
	fmt.Println("concreteComponentA do operation1")
}

func(c *concreteComponentB) operation1() {
	fmt.Println("concreteComponentB do operation1")
}

//装饰器的实现类，每个装饰器实现类会扩展不同的功能
type decoratorWithOperation2 struct {
	*decorator
}
type decoratorWithOperation3 struct {
	*decorator
}

//装饰器实现类对外暴露的接口
func(d *decoratorWithOperation2) operation(){
	d.operation1()
	d.operation2()
}
//装饰器实现类扩展的功能
func (d *decoratorWithOperation2) operation2() {
	fmt.Println("decoratorWithOperation2 do operation2")
}

func(d *decoratorWithOperation3) operation(){
	d.operation1()
	d.operation3()
}

func (d *decoratorWithOperation3) operation3() {
	fmt.Println("decoratorWithOperation3 do operation3")
}