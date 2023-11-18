// 单一职责原则
// 类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个
package main

import "fmt"

/*
// 穿衣服的方式

type Clothes struct {}

// 工作的穿衣方式

func (c *Clothes) Onwork() {
	fmt.Println("工作的装扮.....")
}

// 逛街的穿衣方式

func (c *Clothes) Onshop() {
	fmt.Println("工作的装扮.....")
}

*/

// 逛街类

type ClothesShop struct {
}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

// 工作类

type ClothesWork struct {
}

func (cs *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {

	/*
		c := Clothes{}
		// 工作
		fmt.Println("在工作")
		c.Onwork()
		// 逛街
		fmt.Println("在逛街")
		c.Onshop()
	*/
	// 工作
	cw := ClothesWork{}
	cw.Style()
	// 逛街

	cs := ClothesShop{}
	cs.Style()

}
