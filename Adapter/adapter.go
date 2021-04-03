package Adapter
// 意图：将一个类的接口转换成客户希望的另外一个接口
// 适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
//主要解决：主要解决在软件系统中，常常要将一些"现存的对象"放到新的环境中，而新环境要求的接口是现对象不能满足的。
//优点： 1、可以让任何两个没有关联的类一起运行。
//2、提高了类的复用。
//3、增加了类的透明度。
//4、灵活性好。
//缺点： 1、过多地使用适配器，会让系统非常零乱，不易整体进行把握。比如，明明看到调用的是 A 接口，其实内部被适配成了 B 接口的实现，
//一个系统如果太多出现这种情况，无异于一场灾难。
//因此如果不是很有必要，可以不使用适配器，而是直接对系统进行重构。


//客户希望调用的接口
type Target interface {
	use() string
}
//被适配者
type Adaptee interface {
	specialPlay() string
}
type AdapteeA struct {
	special string
}

type AdapteeB struct {
	special string
}

func(a *AdapteeA) specialPlay() string {
	return a.special
}

func(a *AdapteeB) specialPlay() string {
	return a.special
}

type objectAdapter struct {
	adaptee Adaptee
}
func (o *objectAdapter) use() string{
	return o.adaptee.specialPlay()
}

// 传入一个被适配者，返回一个target实现类，这个类通过实现use方法，
// 来达到target想要的结果
func createTarget(a Adaptee) Target{
	return &objectAdapter{a}
}