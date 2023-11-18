package main

import "fmt"

// 适配的目标

type V5 interface {
	Use5V()
}

// 被适配的角色，适配者

type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("使用220V电压")
}

// 适配器

type Addptoer struct {
	v220 *V220
}

func (a *Addptoer) Use5V() {
	fmt.Println("使用适配器充电.....")
	// 调用适配者方法
	a.v220.Use220V()
}

func NewAddptoer(v220 *V220) *Addptoer {
	return &Addptoer{v220}
}

// 业务类

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

// 普通业务功能

func (p *Phone) Charge() {
	fmt.Println("Phone 进行了充电")
	p.v.Use5V()
}

func main() {
	phone := NewPhone(NewAddptoer(new(V220)))
	phone.Charge()
}
