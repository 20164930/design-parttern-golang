package factory_abstract

//抽象工厂是更高级的抽象，不仅把一个工厂下的产品进行抽象和提取，还把各种工厂抽象，
// 使用一个超级工厂来创建各个子工厂，再用子工厂分别创建各类产品
//优点：当一个产品族中的多个对象被设计成一起工作时，它能保证客户端始终只使用同一个产品族中的对象。
//缺点：产品族扩展非常困难，要增加一个系列的某一产品，既要在抽象的 Creator 里加代码，又要在具体的里面加代码。


//超级工厂类，用来生产各种工厂
type FactoryCreator struct {}

//工厂接口，一种工厂能生产一个产品族
type Factory interface {
	GetFruitProduct(string) Fruit
	GetVegetableProduct(string) Vegetable
}
//水果工厂和蔬菜工厂
type FruitFactory struct {}
type VegetableFactory struct {}


type Fruit interface {
	GetFruitName() string
}
type Vegetable interface {
	GetVegetableQuantity() int
}
//超级工厂方法，用于生产工厂
func (f *FactoryCreator) ProducerFactory(factoryName string) Factory {
	switch factoryName {
	case "fruit":
		return &FruitFactory{}
	case "vegetable":
		return &VegetableFactory{}
	}
	return nil
}

func (f *FruitFactory) GetFruitProduct(name string) Fruit {
	switch name {
	case "apple":
		return &Apple{}
	case "banana":
		return &Banana{}
	}
	return nil
}
//由于接口需要实现所有方法，所以会有多余的无用方法
func (f *FruitFactory) GetVegetableProduct(s string) Vegetable {
	return nil
}

func (v *VegetableFactory) GetVegetableProduct(name string) Vegetable {
	switch name {
	case "tomato":
		return &Tomato{}
	case "potato":
		return &Potato{}
	}
	return nil
}
//由于接口需要实现所有方法，所以会有多余的无用方法
func (v *VegetableFactory) GetFruitProduct(name string) Fruit {
	return nil
}

type Apple struct {
}
type Banana struct {
}

type Tomato struct {
}
type Potato struct {
}
func (a *Apple) GetFruitName() string {
	return "Apple"
}
func (b *Banana) GetFruitName() string {
	return "Banana"
}
func (t *Tomato) GetVegetableQuantity() int {
	return 1
}
func (p *Potato) GetVegetableQuantity() int {
	return 2
}
