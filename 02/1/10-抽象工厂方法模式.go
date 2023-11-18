package main

import "fmt"

// ===== 抽象层 =====

// 抽象的苹果类

type AbstractApple interface {
	ShowApple()
}

// 抽象的香蕉类

type AbstractBanana interface {
	ShowBanana()
}

// 抽象的梨类

type AbstractPear interface {
	ShowPear()
}

// 抽象的工厂

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// ===== 实现层 =====

// 中国产品族

type ChinaApple struct {
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国的苹果")
}

type ChinaBanana struct {
}

func (ca *ChinaBanana) ShowBanana() {
	fmt.Println("中国的香蕉")
}

type ChinaPear struct {
}

func (ca *ChinaPear) ShowPear() {
	fmt.Println("中国的梨")
}

// 中国的抽象工厂

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

/* 日本产品族*/

type JanPanApple struct {
}

func (ja *JanPanApple) ShowApple() {
	fmt.Println("这是日本的苹果")
}

type JanPanBanana struct {
}

func (jb *JanPanBanana) ShowBanana() {
	fmt.Println("这是日本的香蕉")
}

type JanpanPear struct {
}

func (jp *JanpanPear) ShowPear() {
	fmt.Println("这是日本的梨")
}

// 日本的抽象工厂

type JanpanFactory struct {
}

func (jf *JanpanFactory) CreaterApple() AbstractApple {
	var apple AbstractApple
	apple = new(JanPanApple)
	return apple
}

func (jf *JanpanFactory) CreaterBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(JanPanBanana)
	return banana
}

func (jf *JanpanFactory) CreaterPear() AbstractPear {
	var pear AbstractPear
	pear = new(JanpanPear)
	return pear
}

// 针对产品族进行添加，是符合开闭原则的
// 针对产品等级结构进行添加，是不符合开闭原则的

// ===== 逻辑层 =====
func main() {
	// 需求1：需要中国的苹果、香蕉、梨
	// 1. 创建一个中国工厂
	var cFac AbstractFactory
	cFac = new(ChinaFactory)
	// 2. 生成水果
	var cApple AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()
	// 3. 生成香蕉
	var cBanana AbstractBanana
	cBanana = cFac.CreateBanana()
	cBanana.ShowBanana()
	// 4. 生成梨
	var CPear AbstractPear
	CPear = cFac.CreatePear()
	CPear.ShowPear()
}
