package main

//
//// 核心模块 Doctor
//
//type Doctor struct {
//}
//
//func (d *Doctor) treatEye() {
//	fmt.Println("医生治疗眼睛")
//}
//
//func (d *Doctor) treatNose() {
//	fmt.Println("医生治疗鼻子")
//}
//
//// 抽象的命令
//type Command interface {
//	Treat()
//}
//
//// 治疗眼睛的病历单
//
//type CommandThreatEye struct {
//	doctor *Doctor
//}
//
//func (cmd *CommandThreatEye) Treat() {
//	cmd.doctor.treatEye()
//}
//
//// 治疗鼻子的病历单
//
//type CommandThreatNose struct {
//	doctor *Doctor
//}
//
//func (cmd *CommandThreatNose) Treat() {
//	cmd.doctor.treatNose()
//}
//
//// 护石，命令的调用者
//
//type Nurse struct {
//	CmdList []Command // 收集的命令集合
//}
//
//// 发送命令的方法
//
//func (n *Nurse) Notify() {
//	if n.CmdList == nil {
//		return
//	}
//	for _, cmd := range n.CmdList {
//		cmd.Treat() // 多态现象，调用具体命令
//	}
//}
//
//// 病人业务逻辑层
//
//func main() {
//	// 依赖病历单，通过填写病历单，让医生看病
//	doctor := new(Doctor)
//	cmdEye := CommandThreatEye{doctor}
//	//cmdEye.Treat()
//	cmdNose := CommandThreatNose{doctor}
//	//cmdNose.Treat()
//	// 护士
//	nurse := new(Nurse)
//	// 收集病历单
//	nurse.CmdList = append(nurse.CmdList, &cmdEye)
//	nurse.CmdList = append(nurse.CmdList, &cmdNose)
//	// 通知下发指令
//	nurse.Notify()
//}
