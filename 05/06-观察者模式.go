package main

//import "fmt"
//
//// 抽象层
//
//// 抽象的观察者
//type Listener interface {
//	OnTeacherComming() // 观察者得到通知后要触发的动作
//}
//
//type Notifier interface {
//	AddListener(listener Listener)
//	RemoveListener(listener Listener)
//	Notify()
//}
//
//// 实现层
//
//// 观察者具体的学生
//
//type StuZhang3 struct {
//	Badthing string
//}
//
//func (s *StuZhang3) OnTeacherComming() {
//	fmt.Println("张三停止了", s.Badthing)
//}
//
//func (s *StuZhang3) DoBadthing() {
//	fmt.Println("张三正在", s.Badthing)
//}
//
//type StuZhao4 struct {
//	Badthing string
//}
//
//func (s *StuZhao4) OnTeacherComming() {
//	fmt.Println("赵四停止了", s.Badthing)
//}
//
//func (s *StuZhao4) DoBadthing() {
//	fmt.Println("赵四正在", s.Badthing)
//}
//
//type StuWang5 struct {
//	Badthing string
//}
//
//func (s *StuWang5) OnTeacherComming() {
//	fmt.Println("王五停止了", s.Badthing)
//}
//
//func (s *StuWang5) DoBadthing() {
//	fmt.Println("王五正在", s.Badthing)
//}
//
//// 统治者班长
//
//type ClassMonitor struct {
//	listenerList []Listener // 需要通知全部的观察者集合
//}
//
//func (m *ClassMonitor) AddListener(listener Listener) {
//	m.listenerList = append(m.listenerList, listener)
//}
//
//func (m *ClassMonitor) RemoveListener(listener Listener) {
//	for index, l := range m.listenerList {
//		// 找到要删除的元素位置
//		if listener == l {
//			// 将删除的元素前后位置链接起来
//			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
//			break
//		}
//	}
//}
//
//func (m *ClassMonitor) Notify() {
//	for _, listener := range m.listenerList {
//		listener.OnTeacherComming() // 多态现象
//	}
//}
//
//// 业务逻辑层
//func main() {
//	s1 := &StuZhang3{
//		Badthing: "抄作业",
//	}
//	s2 := &StuZhao4{
//		Badthing: "玩王者荣耀",
//	}
//	s3 := &StuWang5{
//		Badthing: "看赵四玩王者荣耀",
//	}
//	classMonitor := new(ClassMonitor)
//	classMonitor.AddListener(s1)
//	classMonitor.AddListener(s2)
//	classMonitor.AddListener(s3)
//	fmt.Println("上课了，但是老师还没有来，学生们都在忙自己的事")
//	s1.DoBadthing()
//	s2.DoBadthing()
//	s3.DoBadthing()
//	fmt.Println("老师来了，班长给学生们使了个眼神儿....")
//	classMonitor.Notify()
//
//}
