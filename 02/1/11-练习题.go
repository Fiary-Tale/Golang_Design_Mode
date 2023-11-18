/*
练习：
	设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
	现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
	要求组装两台电脑，
			    1台（Intel的CPU，Intel的显卡，Intel的内存）
			    1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	用抽象工厂模式实现。
*/

package main

import "fmt"

/*   抽象层    */

type AbstractCpu interface {
	ShowCpu()
}

type AbstractRam interface {
	ShowRam()
}

type AbstractVga interface {
	ShowVga()
}

// 抽象的工厂

type AbstractPcFactory interface {
	CreateCpu() AbstractCpu
	CreateRam() AbstractRam
	CreateVga() AbstractVga
}

/*   实现层    */

// Inter 产品族

type InterCpu struct {
}

func (ic *InterCpu) ShowCpu() {
	fmt.Println("InterCPU开始计算")
}

type InterRam struct {
}

func (ir *InterRam) ShowRam() {
	fmt.Println("InterRam开始存储")
}

type InterVga struct {
}

func (iv *InterVga) ShowVga() {
	fmt.Println("InterVga开始显示")
}

// Inter的抽象工厂

type InterFactory struct {
}

func (If *InterFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(InterCpu)
	return cpu
}

func (If *InterFactory) CreateRam() AbstractRam {
	var ram AbstractRam
	ram = new(InterRam)
	return ram
}
func (If *InterFactory) CreateVga() AbstractVga {
	var vga AbstractVga
	vga = new(InterVga)
	return vga
}

// Nvidia的产品族

type NvidiaCpu struct {
}

func (nc *NvidiaCpu) ShowCpu() {
	fmt.Println("Nvidia_Cpu开始计算")
}

type NvidiaRam struct {
}

func (nc *NvidiaRam) ShowRam() {
	fmt.Println("Nvidia_Ram开始存储")
}

type NvidiaVga struct {
}

func (nc *NvidiaVga) ShowVga() {
	fmt.Println("Nvidia_Vga开始显示")
}

// Nvidia的抽象工厂

type NvidiaFactory struct {
}

func (nf *NvidiaFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(NvidiaCpu)
	return cpu
}

func (nf *NvidiaFactory) CreateRam() AbstractRam {
	var ram AbstractRam
	ram = new(NvidiaRam)
	return ram
}

func (nf *NvidiaFactory) CreateVga() AbstractVga {
	var vga AbstractVga
	vga = new(NvidiaVga)
	return vga
}

// Kingston的产品族

type kingstonCpu struct {
}

func (kc *kingstonCpu) ShowCpu() {
	fmt.Println("kingston_cpu开始计算")
}

type kingstonRam struct {
}

func (kc *kingstonRam) ShowRam() {
	fmt.Println("kingston_ram开始存储")
}

type kingstonVga struct {
}

func (kc *kingstonVga) ShowVga() {
	fmt.Println("kingston_vga开始显示")
}

// kingston的抽象工厂
type kingstonFactory struct {
}

func (kf *kingstonFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(kingstonCpu)
	return cpu
}
func (kf *kingstonFactory) CreateRam() AbstractRam {
	var ram AbstractRam
	ram = new(kingstonRam)
	return ram
}
func (kf *kingstonFactory) CreateVga() AbstractVga {
	var vga AbstractVga
	vga = new(kingstonVga)
	return vga
}

/*   业务逻辑层    */

func main() {
	fmt.Println("需求一：Intel的CPU，Intel的显卡，Intel的内存")
	// 需求1： 1台（Intel的CPU，Intel的显卡，Intel的内存）
	// 1.1 创建一个Inter工厂
	var iFac AbstractPcFactory
	iFac = new(InterFactory)
	// 1.2 生产inter的cpu
	var icpu AbstractCpu
	icpu = iFac.CreateCpu()
	icpu.ShowCpu()
	// 1.3 生产inter的Ram
	var iram AbstractRam
	iram = iFac.CreateRam()
	iram.ShowRam()
	// 1.4 生成inter的Vga
	var ivga AbstractVga
	ivga = iFac.CreateVga()
	ivga.ShowVga()
	fmt.Println("需求二：Intel的CPU， nvidia的显卡，Kingston的内存")
	// 需求2： 1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	// 上述已经创建了Inter的CPU，无需再次创建
	icpu.ShowCpu()
	// 2.1 创建一个nvidia的工厂
	var nFac AbstractPcFactory
	nFac = new(NvidiaFactory)
	// 2.2 生产nvidia的显卡
	var nVga AbstractVga
	nVga = nFac.CreateVga()
	nVga.ShowVga()
	// 2.3 创建一个Kingston的工厂
	var kFac AbstractPcFactory
	kFac = new(kingstonFactory)
	// 2.4 生产kingston的内存
	var kRam AbstractRam
	kRam = kFac.CreateRam()
	kRam.ShowRam()
}
