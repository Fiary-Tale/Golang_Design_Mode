package main

//
///*
//三个要点：
//	一是某个类只能有一个实例；
//	二是它必须自行创建这个实例；
//	三是它必须自行向整个系统提供这个实例
//*/
//
///*
//	保证一个类永远只能有一个对象
//*/
//
//// 1.保证这个类非公优化，外界不能通过这个类创建一个对象
//// 那么这个类应该变得非共有访问，类的名称首字母要小写
//type singelton struct {
//}
//
//// 2.保证一个指针可以指向这个唯一对象，单这个指针永远不能改变方向
//// Golang没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
//
//var instance *singelton = new(singelton)
//
//// 3. 如果全部都是私有化，那么外部模块将永远无法访问到这个对象
//// 所需要对外提供一个方法来获取到这个对象
//
//// GetInstance不可以定义为singelton一个成员方法
//
//func GetInstance() *singelton {
//	return instance
//}
//
//func (s *singelton) SomThing() {
//	fmt.Println("单例某方法")
//}
//
//func main() {
//	s1 := GetInstance()
//	s1.SomThing()
//	s2 := GetInstance()
//	if s1 == s2 {
//		fmt.Println("s1=s2")
//	}
//
//}
