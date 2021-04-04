package facade

import "strconv"

//意图：为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
//主要解决：降低访问复杂系统的内部子系统时的复杂度，简化客户端与之的接口。
//如何解决：客户端不与系统耦合，外观类与系统耦合。
//优点
//1.降低了子系统与客户端之间的耦合度，使得子系统的变化不会影响调用它的客户类。
//2.对客户屏蔽了子系统组件，减少了客户处理的对象数目，并使得子系统使用起来更加容易。
//3.降低了大型软件系统中的编译依赖性，简化了系统在不同平台之间的移植过程，因为编译一个子系统不会影响其他的子系统，也不会影响外观对象。
//
//缺点:
//1.不能很好地限制客户使用子系统类，很容易带来未知风险。
//2.增加新的子系统可能需要修改外观类或客户端的源代码，违背了“开闭原则”。

type fruit interface {
	getName() string
	getPrice() int
}

type facade struct{
	a *apple
	b *banana
}
type apple struct{
	name string
	price int
}

type banana struct{
	name string
	price int
}
func(a *apple) getName() string {
	return a.name
}
func(a *apple) getPrice() int {
	return a.price
}
func(b *banana) getName() string {
	return b.name
}
func(b *banana) getPrice() int {
	return b.price
}
func (f *facade) getAppleAttribute() string {
	return "name: " + f.a.name + ",price: " + strconv.Itoa(f.a.getPrice())
}
func (f *facade) getBananaAttribute() string {
	return "name: " + f.b.name + ",price: " + strconv.Itoa(f.b.getPrice())
}
func createFacade() facade {
	a := &apple{"apple",1}
	b := &banana{"banana",2}
	return facade{a,b}
}