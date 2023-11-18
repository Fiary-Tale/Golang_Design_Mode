package main

import "fmt"

/*
联想路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员MM。根据命令模式，设计烤串场景。
*/

// 核心模块，烤串师傅

type Skewer struct {
}

func (s *Skewer) roasted_lamb() {
	fmt.Println("烤串师傅正在烤羊肉")
}

func (s *Skewer) roasted_chicken() {
	fmt.Println("烤串师傅正在烤鸡翅")
}

// 抽象的命令

type Command interface {
	Roasted()
}

// 烤羊肉的单子

type CommandRoasted_Lamb struct {
	skewer *Skewer
}

func (crl *CommandRoasted_Lamb) Roasted() {
	crl.skewer.roasted_lamb()
}

type CommandRoasted_Chicken struct {
	skewer *Skewer
}

func (crl *CommandRoasted_Chicken) Roasted() {
	crl.skewer.roasted_chicken()
}

// 服务员，分发给烤串师傅

type Nurse struct {
	CmdList []Command
}

// 发送命令的方法

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Roasted() // 多态现象，调用实际命令
	}
}

func main() {
	// 烤串师傅
	skewer := new(Skewer)
	eat_lamb := CommandRoasted_Lamb{skewer}
	eat_chicken := CommandRoasted_Chicken{skewer}
	// 服务员
	nurse := new(Nurse)
	// 收集顾客吃什么
	nurse.CmdList = append(nurse.CmdList, &eat_lamb)
	nurse.CmdList = append(nurse.CmdList, &eat_chicken)
	// 告诉烤串师傅
	nurse.Notify()
}
