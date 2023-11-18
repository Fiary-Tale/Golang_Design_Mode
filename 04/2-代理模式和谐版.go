package main

import "fmt"

//	抽象主题

type BeautyMoman interface {
	// 对男人抛媚眼
	MakeEyesWhithMan()
	// 和男人共度美好时光
	HappyWhithMan()
}

// 具体的主题

type PanJinlian struct {
}

// 对男人抛媚眼

func (p *PanJinlian) MakeEyesWhithMan() {
	fmt.Println("潘金莲对西门庆抛媚眼")
}

// 和男人共度美好时光

func (p *PanJinlian) HappyWhithMan() {
	fmt.Println("潘金莲与西门庆共度良宵")
}

// 中间代理人王婆

type WangPo struct {
	woman BeautyMoman
}

func NewProxy(woman BeautyMoman) BeautyMoman {
	return &WangPo{woman}
}

// 对男人抛媚眼

func (w *WangPo) MakeEyesWhithMan() {
	w.woman.MakeEyesWhithMan()
}

// 和男人共度美好时光

func (w *WangPo) HappyWhithMan() {
	w.woman.HappyWhithMan()

}

// 业务逻辑

func main() {
	// 西门庆想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinlian))
	// 王婆命令金莲抛媚眼
	wangpo.MakeEyesWhithMan()
	// 王婆命令金莲跟西门庆共度良宵
	wangpo.HappyWhithMan()

}
