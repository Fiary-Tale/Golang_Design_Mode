package main

//// ------抽象层------
//
//type Fruit interface {
//	Show() // 接口的方法
//}
//
//// ------实现层------
//
//// 苹果
//
//type Apple struct {
//	Fruit // 继承Fruit类
//}
//
//func (apple *Apple) Show() {
//	fmt.Println("我是苹果")
//}
//
//type Banana struct {
//	Fruit // 继承Fruit类
//}
//
//func (banana *Banana) Show() {
//	fmt.Println("我是香蕉")
//}
//
//type Pear struct {
//	Fruit // 继承Fruit类
//}
//
//func (pear *Pear) Show() {
//	fmt.Println("我是梨")
//}
//
//// ------工厂模块------
//
//type Factory struct {
//}
//
//func (fac *Factory) CreateFruit(kind string) Fruit {
//	var fruit Fruit
//	if kind == "apple" {
//		// apple构造初始化业务
//		fruit = new(Apple) // 满足多态条件的赋值，父类指针指向子类对象
//	} else if kind == "banana" {
//		fruit = new(Banana)
//	} else if kind == "pear" {
//		fruit = new(Pear)
//	}
//	return fruit
//}
//
//// ------业务逻辑层------
//
//func main() {
//	factory := new(Factory)
//
//	apple := factory.CreateFruit("apple")
//	apple.Show()
//
//	banana := factory.CreateFruit("banana")
//	banana.Show()
//
//	pear := factory.CreateFruit("pear")
//	pear.Show()
//}
