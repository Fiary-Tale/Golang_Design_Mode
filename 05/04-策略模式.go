package main

//import "fmt"
//
//// 武器策略(抽象的策略 )
//
//type WeaponStrategy interface {
//	UseWeapon() // 使用武器
//}
//
//// 具体策略AK47武器
//
//type Ak47 struct {
//}
//
//func (ak *Ak47) UseWeapon() {
//	fmt.Println("使用AK47去战斗")
//}
//
//// 具体策略匕首武器
//
//type Knife struct {
//}
//
//func (k *Knife) UseWeapon() {
//	fmt.Println("使用匕首去战斗")
//}
//
//// 环境类
//
//type Hero struct {
//	strategy WeaponStrategy // 拥有一个抽象的策略
//}
//
//// 设置策略方法
//
//func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
//	h.strategy = s
//}
//
//// 战斗方法(业务)
//
//func (h *Hero) Fight() {
//	h.strategy.UseWeapon() // 调用策略
//}
//
//func main() {
//	hero := Hero{}
//	// 更换策略1
//	hero.SetWeaponStrategy(new(Ak47))
//	hero.Fight()
//	// 更换策略2
//	hero.SetWeaponStrategy(new(Knife))
//	hero.Fight()
//}
