package Flyweight
//意图：运用共享技术有效地支持大量细粒度的对象。
//主要解决：在有大量对象时，有可能会造成内存溢出，我们把其中共同的部分抽象出来，
//如果有相同的业务请求，直接返回在内存中已有的对象，避免重新创建。

//优点：
//1.大大减少对象的创建，降低系统的内存，使效率提高。
//缺点：
//1.提高了系统的复杂度，需要分离出外部状态和内部状态，而且外部状态具有固有化的性质,不应该随着内部状态的变化而变化，否则会造成系统的混乱。

//享元抽象类
type flyweight interface{
	operation() string
}
//外部角色
type outFlyweight struct {
	info string
}

func (out *outFlyweight) setInfo (info string) {
	out.info = info
}
func (out *outFlyweight) getInfo() string {
	return out.info
}
//享元实现类,必须包含内部角色和外部角色，内部角色是固定的，外部是变化的
type concreteFlyweight struct {
	key string
	out outFlyweight
}
func (c *concreteFlyweight) operation() string {
	return "flyweight-" + c.key + ",outFlyweight-" + c.out.getInfo()
}

type flyweightFactory struct {
	flyweightMap map[string]flyweight
}

func (f *flyweightFactory) getFlyweight(key string) flyweight {
	if f.flyweightMap == nil {
		f.flyweightMap = make(map[string]flyweight)
	}
	fw, ok := f.flyweightMap[key]
	if !ok {
		concrete := concreteFlyweight{key, outFlyweight{key}}
		f.flyweightMap[key] = &concrete
		fw = &concrete
	}
	return fw
}