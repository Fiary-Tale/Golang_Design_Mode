package main

import "fmt"

// --------> 抽象层 <-------------

type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

// --------> 实现层 <-------------

// 奔驰

type Benz struct {
	//....
}

func (b *Benz) Run() {
	fmt.Println("奔驰已启动.....")
}

// 宝马

type BMW struct {
	//....
}

func (b *BMW) Run() {
	fmt.Println("宝马已启动.....")
}

// 张三司机

type zhangsan struct {
	// ....
}

func (z3 *zhangsan) Driver(car Car) {
	fmt.Println("张三开车")
	car.Run()
}

// 李四司机

type lisi struct {
	// ....
}

func (l4 *lisi) Driver(car Car) {
	fmt.Println("李四开车")
	car.Run()
}

// 王五司机

type wangwu struct {
	// ....
}

func (w5 *wangwu) Driver(car Car) {
	fmt.Println("王五开车")
	car.Run()
}

// --------> 业务逻辑层 <-------------

func main() {
	// 张三开奔驰
	var benz Car
	benz = new(Benz)

	var zhang3 Driver
	zhang3 = new(zhangsan)
	zhang3.Driver(benz)
	// 李四开宝马
	var bmw Car
	bmw = new(BMW)

	var li4 Driver
	li4 = new(lisi)
	li4.Driver(bmw)

	// 王五开奔驰
	var w5 Driver
	w5 = new(wangwu)
	w5.Driver(benz)

}
