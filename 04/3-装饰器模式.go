package main

//
//// 抽象层
//
//type Phone interface {
//	Show() // 构建功能
//}
//
//// 抽象的装饰类，装饰器的基础类(该类本应该是interface，但是golang的interface语法不可以有成员属性)
//
//type Decorator struct {
//	phone Phone // 组合进来抽象的phone
//}
//
//func (d *Decorator) Show() {
//
//}
//
//// 实现层
//
//type Huawei struct {
//}
//
//func (hw *Huawei) Show() {
//	fmt.Println("秀出了HuaWei手机")
//}
//
//type XiaoMi struct {
//}
//
//func (xm *XiaoMi) Show() {
//	fmt.Println("秀出了XiaoMi手机")
//}
//
//// 具体的装饰器	贴膜
//
//type MoDecorator struct {
//	Decorator // 继承基础的装饰器类(主要继承Phone的成员属性)
//}
//
//func (md *MoDecorator) Show() {
//	md.phone.Show() // 调用被装饰的构建的原方法
//
//	fmt.Println("贴膜的手机")
//}
//func NewMoDecorator(phone Phone) Phone {
//	return &MoDecorator{Decorator{phone}}
//}
//
//// 具体的装饰器   手机壳
//
//type KeDecorator struct {
//	Decorator
//}
//
//func (kd *KeDecorator) Show() {
//	kd.phone.Show() // 继承基础的装饰器类(主要继承Phone的成员属性)
//	fmt.Println("带壳的手机")
//}
//
//func NewKeDecorator(phone Phone) Phone {
//	return &KeDecorator{Decorator{phone}}
//}
//
//// 逻辑层
//
//func main() {
//	var huawei Phone
//	huawei = new(Huawei)
//	huawei.Show()
//	fmt.Println("============================")
//	// 贴膜HuaWei手机
//	var moHuaWei Phone
//	moHuaWei = NewMoDecorator(huawei)
//	moHuaWei.Show()
//
//	fmt.Println("============================")
//	// 带壳的HuaWei手机
//	var keHuaWei Phone
//	keHuaWei = NewKeDecorator(huawei)
//	keHuaWei.Show()
//	fmt.Println("============================")
//	// 带壳贴膜的华为手机		贴膜--> 带壳
//	var MoKeHuaWei Phone
//	MoKeHuaWei = NewKeDecorator(moHuaWei)
//	MoKeHuaWei.Show()
//	fmt.Println("============================")
//	// 带壳贴膜的华为手机		带壳--> 贴膜
//	var KeMoHuaWei Phone
//	KeMoHuaWei = NewMoDecorator(keHuaWei)
//	KeMoHuaWei.Show()
//
//}
