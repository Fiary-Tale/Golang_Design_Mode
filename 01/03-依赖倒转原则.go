// 依赖于抽象(接口),不要依赖具体的实现(类),也就是针对接口编程
package main

//import "fmt"
//
//// 1. 耦合度极高的模块设计
//
//// 现有张三、李四、宝马、奔驰
//// 1.1张三开奔驰
//// 1.2李四开宝马
//// 奔驰汽车
//
//type Benz struct {
//}
//
//func (b *Benz) Run() {
//	fmt.Println("奔驰已启动......")
//}
//
//// 宝马汽车
//
//type BMW struct {
//}
//
//func (b *BMW) Run() {
//	fmt.Println("宝马已启动......")
//}
//
//// 司机张三
//
//type Zhang3 struct {
//}
//
//// 张三开奔驰的能力
//
//func (z3 *Zhang3) DriveBenz(benz *Benz) {
//	fmt.Println("张三开奔驰")
//	benz.Run()
//}
//
//// 张三开宝马的能力(+)
//
//func (z3 *Zhang3) DriveBmw(bmw *BMW) {
//	fmt.Println("张三开宝马")
//	bmw.Run()
//}
//
//// 司机李四
//
//type Li4 struct {
//}
//
//// 李四开宝马的能力
//
//func (l4 *Li4) DriveBMW(bmw *BMW) {
//	fmt.Println("李四开宝马")
//	bmw.Run()
//}
//
//// 李四开奔驰的能力 +
//
//func (l4 *Li4) DriveBenz(benz *Benz) {
//	fmt.Println("李四开奔驰")
//	benz.Run()
//}
//
//func main() {
//	// 1.1张三开奔驰
//	benz := &Benz{}
//	zhang3 := &Zhang3{}
//	zhang3.DriveBenz(benz)
//
//	// 1.2李四开宝马
//	bmw := &BMW{}
//	l4 := &Li4{}
//	l4.DriveBMW(bmw)
//
//	// 现在张三要开宝马 +
//	zhang3.DriveBmw(bmw)
//	// 现在李四要开奔驰 +
//	l4.DriveBenz(benz)
//}
