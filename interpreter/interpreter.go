package interpreter

import (
	"fmt"
	"strings"
)

//主要解决：对于一些固定文法构建一个解释句子的解释器。
//何时使用：如果一种特定类型的问题发生的频率足够高，那么可能就值得将该问题的各个实例表述为一个简单语言中的句子。
//这样就可以构建一个解释器，该解释器通过解释这些句子来解决该问题。

//如何解决：构建语法树，定义终结符与非终结符。
//优点：
//1、可扩展性比较好，灵活。
//2、增加了新的解释表达式的方式。
//3、易于实现简单文法。
//
//缺点：
//1、可利用场景比较少。
//2、对于复杂的文法比较难维护。
//3、解释器模式会引起类膨胀。
//4、解释器模式采用递归调用方法。

//抽象表达式
type Expression interface {
	interpret(string) bool
}
//环境类，主要用来传递共享变量
type Context struct {
	sex []string
	isMarried []string
}
//终结表达式，表示最终需要做判断的词汇或问题
type terminalExpression struct{
	data map[string]bool
}
func (t *terminalExpression) initTerminal(infos []string) {
	if t.data == nil {
		t.data = make(map[string]bool)
	}
	for _,v := range infos {
		t.data[v] = true
	}
}

//非终结表达式，可以有多个，通常会递归调用，并且会调用周围别的非终结表达式
type andExpression struct {
	expr1 Expression
	expr2 Expression
}

var context Context
var sexAndMarried Expression

func init() {
	sex := []string{"man", "woman"}
	isMarried := []string{"married"}
	context = Context{isMarried, sex}
	sexExpression :=&terminalExpression{}
	sexExpression.initTerminal(sex)

	marriedExpression := &terminalExpression{}
	marriedExpression.initTerminal(isMarried)

	sexAndMarried = &andExpression{marriedExpression, sexExpression}
}

func (e *terminalExpression) interpret(info string) bool {
	_, ok := e.data[info]
	return ok
}

func (e *andExpression) interpret(info string) bool {
	infos := strings.Split(info, " ")
	return e.expr1.interpret(infos[0]) && e.expr2.interpret(infos[1])
}

func(c *Context) judge(info string) {
	infos := strings.Split(info, "is a ")
	ret := sexAndMarried.interpret(strings.TrimSpace(infos[1]))
	if ret {
		fmt.Println(info + ", can get the prize")
	} else {
		fmt.Println(info + ", can not get the prize")
	}
}