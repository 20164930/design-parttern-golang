package simplefactory
// 优点： 1、一个调用者想创建一个对象，只要知道其名称就可以了。
// 2、扩展性高，如果想增加一类产品，只要扩展一个工厂的产品就可以。
// 3、屏蔽产品的具体实现，调用者只关心产品的接口。

//缺点：每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，
//在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。这并不是什么好事。
type Product interface {
	getQuality() int
}

type ProductA struct {}
type ProductB struct {}
type ProductC struct {}

func (p *ProductA) getQuality() int{
	return 1
}

func (p *ProductB) getQuality() int{
	return 2
}

func (p *ProductC) getQuality() int{
	return 3
}

type Factory struct {

}

func (f *Factory) Create(productName string) Product{
	switch productName {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	case "C":
		return &ProductC{}
	}
	return nil
}