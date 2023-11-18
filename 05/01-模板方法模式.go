package main

import (
	"fmt"
)

// 抽象类，制作饮料，包含一个模板，全部实现步骤

type Beverage interface {
	// 煮开水
	BoilWater()
	// 冲泡
	Brew()
	// 倒入杯中
	PourInCup()
	// 添加佐料
	AddThings()
	// 是否添加佐料的Hook
	WantAddThings() bool
}

// 封装一套流程模板基础类，让具体的制作流程继承且实现

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	// 固定流程

	t.b.BoilWater() // 真实是执行子类具体的实现方法
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

// 具体的制作饮料的流程，制作咖啡

type MakeCoffee struct {
	template
}

// 煮开水

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

// 冲泡

func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲泡咖啡豆")
}

// 倒入杯中

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将冲泡好的咖啡倒入杯中")
}

// 添加佐料

func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	// b是Beverage是MakeCoffee的接口，这些需要给接口赋值，让b指向具体的子类
	// 来触发b的全部方法的多态性
	makeCoffee.b = makeCoffee
	return makeCoffee
}

// 具体制作茶叶的流程

type MakeTea struct {
	template
}

// 煮开水

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水冲泡到80度")
}

// 冲泡茶叶

func (mt *MakeTea) Brew() {
	fmt.Println("用水冲泡茶叶")
}

// 倒入杯中

func (mt *MakeTea) PourInCup() {
	fmt.Println("将冲泡好的茶叶水倒入茶壶中")
}

// 添加佐料

func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mc *MakeTea) WantAddThings() bool {
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

// 业务逻辑层

func main() {

	fmt.Println("=============================制作一杯咖啡=============================")
	// 1. 制作一杯咖啡
	makecoffee := NewMakeCoffee()
	makecoffee.MakeBeverage()
	fmt.Println("=============================制作一杯茶=============================")
	// 2.制作一杯茶
	maketea := NewMakeTea()
	maketea.MakeBeverage()

}
