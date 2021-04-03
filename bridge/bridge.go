package bridge

// 将抽象与实现分离，使它们可以独立变化。它是用组合关系代替继承关系来实现，
// 从而降低了抽象和实现这两个可变维度的耦合度。
// 桥接（Bridge）模式的优点是：
//1.抽象与实现分离，扩展能力强
//2.符合开闭原则
//3.符合合成复用原则
//4.其实现细节对客户透明

//缺点是：由于聚合关系建立在抽象层，要求开发者针对抽象化进行设计与编程，
//能正确地识别出系统中两个独立变化的维度，这增加了系统的理解与设计难度。

type Color interface {
	getColorName() string
	getColorCode() string
}
type Red struct {
}
type Green struct {
}
func(r *Red) getColorName() string {
	return "Red"
}
func(r *Red) getColorCode() string {
	return "#DC143C"
}

func(g *Green) getColorName() string {
	return "Green"
}
func(g *Green) getColorCode() string {
	return "#00FF00"
}

//此处可有多个变化的维度，例如重量，cpu型号等
type phone struct {
	color Color
	//any other attribute
}

type huaweiPhone struct {
	*phone
	series string
}

func(p *huaweiPhone) getAttribute () string {
	return p.series + "/" + p.color.getColorName() + ":" + p.color.getColorCode()
}
