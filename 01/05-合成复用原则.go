package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 添加猫的睡觉方法(使用继承实现)

type CatB struct {
	Cat
}

func (cb *CatB) sleep() {
	fmt.Println("小猫睡觉")
}

// 添加猫的睡觉方法(使用组合实现)

type CatC struct {
	//C *Cat // 组合进来一个Cat类
}

func (cc *CatC) Eat(c *Cat) {
	//cc.C.Eat()
	c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := &Cat{}
	c.Eat()
	fmt.Println("---------------------------------------")
	cb := &CatB{}
	cb.sleep()
	cb.Eat()
	fmt.Println("---------------------------------------")
	cc := &CatC{}
	cc.Eat(c)
	cc.Sleep()
}
