package iterator

//意图：提供一种方法顺序访问一个聚合对象中各个元素, 而又无须暴露该对象的内部表示。
//主要解决：不同的方式来遍历整个整合对象。

//主要优点如下。
//1.访问一个聚合对象的内容而无须暴露它的内部表示。
//2.遍历任务交由迭代器完成，这简化了聚合类。
//3.它支持以不同方式遍历一个聚合，甚至可以自定义迭代器的子类以支持新的遍历。
//4.增加新的聚合类和迭代器类都很方便，无须修改原有代码。
//5.封装性良好，为遍历不同的聚合结构提供一个统一的接口。
//
//其主要缺点是：
//1.增加了类的个数，这在一定程度上增加了系统的复杂性。

type aggregate interface {
	add(i interface{})
	remove(i interface{})
	getIterator() Iterator
}

type Iterator interface {
	first() interface{}
	next() interface{}
	hasNext() bool
}
type iterator struct {
	list []interface{}
	index int
}

func(i *iterator) first() interface{} {
	if len(i.list) > 0 {
		return i.list[0]
	}
	return nil
}
func (i *iterator) hasNext() bool {
	if len(i.list) == 0 {
		return false
	}
	if len(i.list) > i.index + 1 {
		return true
	}
	return false
}

func (i *iterator) next() interface{} {
	if i.hasNext() {
		i.index = i.index + 1
		return i.list[i.index]
	}
	return nil
}
type aggregateA struct{
	list []interface{}
}
func(a *aggregateA) add(i interface{}) {
	a.list = append(a.list, i)
}
func(a *aggregateA) remove(i interface{}) {
	for k:=0; k<len(a.list); k++ {
		if a.list[k] == i {
			a.list = append(a.list[:k], a.list[k+1:]...)
			k--
		}
	}
}
func(a *aggregateA) getIterator() Iterator {
	return &iterator{a.list, -1}
}