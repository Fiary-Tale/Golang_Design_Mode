package main

import "fmt"

/*
// 抽象的策略

type ShopStrategy interface {
	UseStrategy() // 使用策略
}

// 具体的策略A 商场8折

type Discount struct {
}

func (d *Discount) UseStrategy() {
	fmt.Println("商场打八折")
}

// 具体的策略B 商场满减100

type Full_discount struct {
}

func (fd *Full_discount) UseStrategy() {
	fmt.Println("商场满200减100")
}

// 环境类

type Env struct {
	strategy ShopStrategy // 拥有一个抽象的策略
}

// 设置策略方法

func (e *Env) SetShopStrategy(s ShopStrategy) {
	e.strategy = s
}

// 具体实现方法

func (e *Env) Fight() {
	e.strategy.UseStrategy()
}

func main() {
	env := Env{}
	// 策略A 商场8折
	env.SetShopStrategy(new(Discount))
	env.Fight()
	// 策略B 商场满200减100
	env.SetShopStrategy(new(Full_discount))
	env.Fight()
}
*/

// 销售策略

type SellStrategy interface {
	// 根据原价得到售卖价
	GetPrice(price float64) float64
}

type StrategyA struct {
}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A，商场打八折")
	return price * 0.8
}

type StrategyB struct {
}

func (sb *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B，满200减100")
	if price >= 200 {
		price -= 100
	}
	return price
}

// 环境类

type Goods struct {
	Price    float64
	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值", g.Price, ".")
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	nike := Goods{
		Price: 200.0,
	}
	// 上午执行策略A
	nike.SetStrategy(new(StrategyA))
	fmt.Println("上午耐克鞋卖", nike.SellPrice())
	// 下午执行策略B
	nike.SetStrategy(new(StrategyB))
	fmt.Println("下午耐克鞋卖", nike.SellPrice())
}
