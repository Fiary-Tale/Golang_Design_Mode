package main

//// ------抽象层------
//
//// 水果类(抽象的接口)
//
//type Fruit interface {
//	Show()
//}
//
//// 工厂类(抽象的接口)
//
//type AbstractFactory interface {
//	CreateFruit() Fruit
//}
//
//// ------基础模块层------
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
//// 日本梨 ++
//
//type JanpanPear struct {
//	Fruit // 继承Fruit类
//}
//
//func (jp *JanpanPear) Show() {
//	fmt.Println("我是日本梨")
//}
//
//// ------基础工厂模块-------
//
//// 具体的苹果工厂
//
//type AppleFactory struct {
//	AbstractFactory
//}
//
//func (fac *AppleFactory) CreateFruit() Fruit {
//	var fruit Fruit
//	// 生产一个具体的苹果
//	fruit = new(Apple)
//	return fruit
//}
//
//// 具体的香蕉工厂
//
//type BananaFactory struct {
//	AbstractFactory
//}
//
//func (fac *BananaFactory) CreateFruit() Fruit {
//	var fruit Fruit
//	// 生产一个具体的香蕉
//	fruit = new(Banana)
//	return fruit
//}
//
//// 具体的梨工厂
//
//type PearFactory struct {
//	AbstractFactory
//}
//
//func (fac *PearFactory) CreateFruit() Fruit {
//	var fruit Fruit
//	// 生产一个具体的梨
//	fruit = new(Pear)
//	return fruit
//}
//
//// 具体的日本梨工厂 +
//
//type JanpanPearFactory struct {
//	AbstractFactory
//}
//
//func (fac *JanpanPearFactory) CreateFruit() Fruit {
//	var fruit Fruit
//	// 生产一个具体的日本梨
//	fruit = new(JanpanPear)
//	return fruit
//}
//
//// ------业务逻辑层------
//
//func main() {
//	// 需求1：需要一个具体的苹果对象
//	// 1.先要一个具体的苹果工厂
//	var appleFac AbstractFactory
//	appleFac = new(AppleFactory)
//	// 2.生产一个具体的水果
//	var apple Fruit
//	apple = appleFac.CreateFruit()
//	apple.Show() // 多态
//	// --------------------------------------------------------------------
//	// 需求2：需要一个具体的香蕉对象
//	// 1.先要一个具体的香蕉工厂
//	var bananaFac AbstractFactory
//	bananaFac = new(BananaFactory)
//	// 2.生成一个具体的水果
//	var banana Fruit
//	banana = bananaFac.CreateFruit()
//	banana.Show()
//	// 需求3：需要一个具体的梨对象
//	// 1.先要一个具体的梨工厂
//	var pearFac AbstractFactory
//	pearFac = new(PearFactory)
//	// 2.生成一个具体的水果
//	var pear Fruit
//	pear = pearFac.CreateFruit()
//	pear.Show()
//	// 需求4：需要一个具体的日本梨对象
//	// 1.先要一个具体的日本梨工厂
//	var jpFac AbstractFactory
//	jpFac = new(JanpanPearFactory)
//	// 2.生成一个具体的水果
//	var jppear Fruit
//	jppear = jpFac.CreateFruit()
//	jppear.Show()
//}
