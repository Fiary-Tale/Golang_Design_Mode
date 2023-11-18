package main

//import "fmt"
//
//type Goods struct {
//	King string // 商品的种类
//	Fact bool   // 商品的真伪
//}
//
//// ===============抽象层===============
//type Shopping interface {
//	Buy(goods *Goods) // 某任务
//}
//
//// ===============实现层===============
//
//type KoreaShopping struct {
//}
//
//func (ks *KoreaShopping) Buy(goods *Goods) {
//	fmt.Println("去韩国进行了购物，买了", goods.King)
//}
//
//type AmericasShopping struct {
//}
//
//func (as *AmericasShopping) Buy(goods *Goods) {
//	fmt.Println("去美国进行了购物，买了", goods.King)
//}
//
//type AfrikaShopping struct {
//}
//
//func (as *AfrikaShopping) Buy(goods *Goods) {
//	fmt.Println("去非洲进行了购物，买了", goods.King)
//}
//
////海外代理
//
//type OverSeasProxy struct {
//	shopping Shopping // 代理某个主题，抽象的数据类型
//}
//
//func NewProxy(shopping Shopping) Shopping {
//	return &OverSeasProxy{shopping: shopping}
//}
//
//func (op *OverSeasProxy) Buy(goods *Goods) {
//	// 1. 先辨别真伪
//	if op.distinguish(goods) == true {
//		// 2. 调用具体要被代理的购物方式Buy()
//		op.shopping.Buy(goods) // 调用具体方法
//		// 3. 海关安检
//		op.check(goods)
//	}
//
//}
//
//func (op *OverSeasProxy) distinguish(goods *Goods) bool {
//	fmt.Println("对[", goods.King, "]进行了辨别真伪")
//	if goods.Fact == false {
//		fmt.Println("发现假货", goods.King, ",不应该购买")
//	}
//	return goods.Fact
//}
//
//// 安检流程
//
//func (op *OverSeasProxy) check(goods *Goods) {
//	fmt.Println("对[", goods.King, "]进行了海关检查，成功带回")
//}
//
//// ===============业务逻辑层===============
//
//func main() {
//	g1 := Goods{
//		King: "韩国面膜",
//		Fact: true,
//	}
//
//	g2 := Goods{
//		King: "CET4证书",
//		Fact: false,
//	}
//	var KShoping Shopping
//	KShoping = new(KoreaShopping)
//
//	var proxy Shopping
//	proxy = NewProxy(KShoping)
//	proxy.Buy(&g1)
//
//	var AShoping Shopping
//	AShoping = new(AmericasShopping)
//	proxy = NewProxy(AShoping)
//	proxy.Buy(&g2)
//}
