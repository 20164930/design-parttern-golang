package factory
//简单工厂模式违背了开闭原则，而“工厂方法模式”是对简单工厂模式的进一步抽象化，
//其好处是可以使系统在不修改原来代码的情况下引进新的产品，即满足开闭原则。

//客户只知道创建产品的工厂名，而不知道具体的产品名。如 TCL 电视工厂、海信电视工厂等。
//创建对象的任务由多个具体子工厂中的某一个完成，而抽象工厂只提供创建产品的接口。
//客户不关心创建产品的细节，只关心产品的品牌
type AbstractFactory interface{
	createProduct() Product
}
type FactoryA struct{
	name string
}
type FactoryB struct{
	name string
}
type Product interface{
	show() string
}
type ProductA struct{
	name string
}
type ProductB struct{
	name string
}
func(f *FactoryA) createProduct() Product {
	return &ProductA{"productA"}
}
func(f *FactoryB) createProduct() Product {
	return &ProductB{"productB"}
}
func(p *ProductA) show() string {
	return p.name
}
func(p *ProductB) show() string {
	return p.name
}