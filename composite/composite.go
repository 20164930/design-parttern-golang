package composite

import "fmt"

//意图：将对象组合成树形结构以表示"部分-整体"的层次结构。
//组合模式使得用户对单个对象和组合对象的使用具有一致性。
//
//主要解决：它在我们树型结构的问题中，模糊了简单元素和复杂元素的概念，客户程序可以像处理简单元素一样来处理复杂元素，从
//而使得客户程序与复杂元素的内部结构解耦。
//
//组合模式的主要优点有：
//1.组合模式使得客户端代码可以一致地处理单个对象和组合对象，无须关心自己处理的是单个对象，还是组合对象，这简化了客户端代码；
//2.更容易在组合体内加入新的对象，客户端不会因为加入了新的对象而更改源代码，满足“开闭原则”；
//
//其主要缺点是：
//1.设计较复杂，客户端需要花更多时间理清类之间的层次关系；
//2.不容易限制容器中的构件；
//3.不容易用继承的方法来增加构件的新功能；

type Component interface {
	add(Component)
	remove(Component)
	getChild(int)Component
	printComponent()
}

type employee struct {
	id string
}

type group struct {
	children []Component
}

func (e *employee) add(Component) {}
func (e *employee) remove(Component) {}
func (e *employee) getChild(int) Component{return nil}
func (e *employee) printComponent(){
	fmt.Println(e.id)
}

func (g *group) add(c Component) {
	g.children = append(g.children, c)
}

func (g *group) remove(c Component) {
	for i:=0; i<len(g.children);i++ {
		// 此处自行实现equals方法
		if g.children[i] == c {
			g.children = append(g.children[:i],g.children[i+1:]...)
			i--
		}
	}
}

func (g *group) getChild(i int) Component {
	return g.children[i]
}
func (g *group) printComponent() {
	for i:=0;i<len(g.children);i++ {
		g.children[i].printComponent()
	}
}

