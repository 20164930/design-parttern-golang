package builder

// 建造者模式是创建一个director来使用不同的builder实现类，去组装各种产品，
// 产品往往有多个属性，且有一些属性是必选，有一些是可选
// 具体的builder用来创建属性，director用来组装这些属性

// 该模式的主要优点如下：
//封装性好，构建和表示分离。
//扩展性好，各个具体的建造者相互独立，有利于系统的解耦。
//客户端不必知道产品内部组成的细节，建造者可以对创建过程逐步细化，而不对其它模块产生任何影响，便于控制细节风险。

//其缺点如下：
//产品的组成部分必须相同，这限制了其使用范围。
//如果产品的内部变化复杂，如果产品内部发生变化，则建造者也要同步修改，后期维护成本较大。


//（Product）一顿肯德基，可能包含的内容，其中humBurger是必选，其余是可选
type Meal struct {
	cola int
	chips int
	pizza int
	humBurger int
}

//建造者，可以生产所有的可选属性，并且可以返回一个product
type Builder interface {
	buyCola()
	buyChips()
	buyPizza()
	getProduct() *Meal
}

//指挥者，用来管理整个建造过程
type Director struct {
}
func (d *Director) buyMeal(builder Builder) *Meal{
	builder.buyCola()
	builder.buyChips()
	builder.buyPizza()
	return builder.getProduct()
}

//两种套餐的建造者，可以生产各自负责的套餐（product）
type mealABuilder struct {
	meal *Meal
}
type mealBBuilder struct {
	meal *Meal
}
//建造者的构造函数，可用工厂函数进行生成
func GetABuilder() *mealABuilder {
	return &mealABuilder{&Meal{humBurger: 1}}
}
func GetBBuilder() *mealBBuilder {
	return &mealBBuilder{&Meal{humBurger: 2}}
}
func(m *mealABuilder) buyCola(){
	m.meal.cola = 1
}
func(m *mealABuilder) buyChips(){
	m.meal.chips= 1
}
func(m *mealABuilder) buyPizza(){
	m.meal.pizza= 1
}
func(m *mealABuilder) getProduct() *Meal{
	return m.meal
}


func(m *mealBBuilder) buyCola(){
	m.meal.cola = 1
}
func(m *mealBBuilder) buyChips(){
	m.meal.chips= 2
}
func(m *mealBBuilder) buyPizza(){
	m.meal.pizza= 0
}
func(m *mealBBuilder) getProduct() *Meal{
	return m.meal
}