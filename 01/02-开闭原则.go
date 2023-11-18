// 类的改动是通过增加代码进行的，而不是修改源码，一旦代码落地，最好不要进行修改，是以添加的形式进行新增功能

package main

import "fmt"

/*
// 无开闭原则平铺模式

type Banker struct {
}

// 存款业务

func (b *Banker) Save() {
	fmt.Println("进行了存款业务.....")
}

// 转账业务

func (b *Banker) Tansfer() {
	fmt.Println("进行了转账业务.....")
}

// 支付

func (b *Banker) Pay() {
	fmt.Println("进行了支付业务.....")
}

// 股票业务(+)

func (b *Banker) Shares() {
	fmt.Println("进行了股票业务.....")
}

func Banker01() {
	banker := &Banker{}
	banker.Save()
	banker.Tansfer()
	banker.Pay()
	banker.Shares()
}
*/

// 开闭原则设计

// 抽象的银行业务员

type AbstractBanker interface { // 抽象的处理业务接口
	DoBusi()
}

// 存款业务

type SaveBanker struct {
	// AbstractBanker
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款业务.....")
}

// 支付业务

type PayBanker struct {
	// AbstractBanker
}

func (sb *PayBanker) DoBusi() {
	fmt.Println("进行了支付业务.....")
}

// 转账业务 (+)

type TranferBanker struct {
	// AbstractBanker
}

func (st *TranferBanker) DoBusi() {
	fmt.Println("进行了转账业务....")
}

// 股票业务(+)

type SharesBanker struct {
	// AbstractBanker
}

func (sb *SharesBanker) DoBusi() {
	fmt.Println("进行了股票业务....")
}

// 实现一个架构层(基于抽象层进行业务封装-针对interface接口进行封装)

func BankBusiness(banker AbstractBanker) {
	// 通过接口向下调用(多态的现象)
	banker.DoBusi()
}

func Banker02() {
	//// 存款业务(原有)
	//sb := SaveBanker{}
	//sb.DuBusi()
	//
	//// 转账业务(新增)
	//st := TranferBanker{}
	//st.DuBusi()
	//
	//// 股票业务(新增)
	//sb2 := SharesBanker{}
	//sb2.DuBusi()

	// 存款业务(原有)
	BankBusiness(&SaveBanker{})
	// 转账业务(新增)
	BankBusiness(&TranferBanker{})
	// 股票业务(新增)
	BankBusiness(&SharesBanker{})
}

func main() {
	//Banker01()
	Banker02()
}
