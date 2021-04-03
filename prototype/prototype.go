package prototype

import (
	"encoding/json"
	"reflect"
)

//原型模式（Prototype Pattern）是用于创建重复的对象，同时又能保证性能。
//这种模式是实现了一个原型接口，该接口用于创建当前对象的克隆。当直接创建对象的代价比较大时，则采用这种模式。
//优点： 1、性能提高。 2、逃避构造函数的约束。
//缺点： 1、配备克隆方法需要对类的功能进行通盘考虑，这对于全新的类不是很难，但对于已有的类不一定很容易，
//特别当一个类引用不支持串行化的间接对象，或者引用含有循环结构的时候。
type Manage struct{}

func (m *Manage) create(name string) Product {
	switch name {
	case "A":
		return &ProductA{"A"}
	case "B":
		return &ProductB{"B"}
	}
	return nil
}

type Product interface {
	clone() Product
	getName() string
}

type ProductA struct {
	name string
}
type ProductB struct {
	name string
}

func (p *ProductA) getName() string {
	return p.name
}

func (p *ProductA) clone() Product {
	return &ProductA{p.name}
}

func (p *ProductB) getName() string {
	return p.name
}

func (p *ProductB) clone() Product {
	return &ProductB{p.name}
}


//深度克隆，可以克隆任意数据类型
func DeepClone(src interface{}) interface{} {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr { //如果是指针类型
		typ = typ.Elem()                          //获取源实际类型(否则为指针类型)
		dst := reflect.New(typ).Elem()            //创建对象
		b, _ := json.Marshal(src)                 //导出json
		json.Unmarshal(b, dst.Addr().Interface()) //json序列化
		return dst.Addr().Interface()             //返回指针
	} else {
		dst := reflect.New(typ).Elem()            //创建对象
		b, _ := json.Marshal(src)                 //导出json
		json.Unmarshal(b, dst.Addr().Interface()) //json序列化
		return dst.Interface()                    //返回值
	}
}

//浅克隆，可以克隆任意数据类型，对指针类型子元素无法克隆
//获取类型：如果类型是指针类型，需要使用Elem()获取对象实际类型
//获取实际值：如果值是指针类型，需要使用Elem()获取实际数据
//说白了，Elem()就是获取反射数据的实际类型和实际值
func Clone(src interface{}) interface{} {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr { //如果是指针类型
		typ = typ.Elem()               //获取源实际类型(否则为指针类型)
		dst := reflect.New(typ).Elem() //创建对象
		data := reflect.ValueOf(src)   //源数据值
		data = data.Elem()             //源数据实际值（否则为指针）
		dst.Set(data)                  //设置数据
		dst = dst.Addr()               //创建对象的地址（否则返回值）
		return dst.Interface()         //返回地址
	} else {
		dst := reflect.New(typ).Elem() //创建对象
		data := reflect.ValueOf(src)   //源数据值
		dst.Set(data)                  //设置数据
		return dst.Interface()         //返回
	}
}