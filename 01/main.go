//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_+-="
//
//func randStringBytes(n int) string {
//	b := make([]byte, n)
//	for i := range b {
//		b[i] = letterBytes[rand.Intn(len(letterBytes))]
//	}
//	return string(b)
//}
//
//func main() {
//	rand.Seed(time.Now().UnixNano()) // 设置随机种子
//	fmt.Println(randStringBytes(5))  // 生成长度为5的随机字符串
//	//for i := 0; i < 10; i++ {
//	//	fmt.Println(randStringBytes(5))
//	//}
//	TimeStr := time.Now().Format("01-02 15:04:05")
//	for i := 0; i < 10; i++ {
//		time.Sleep(time.Duration(1) * time.Second)
//		fmt.Println(TimeStr)
//	}
//}

package main

import (
	"fmt"
)

//func Transer(newstr string) string {
//
//	return strconv.QuoteToASCII(newstr)
//
//}
//func main() {
//	//fmt.Println("将要转换的编码中文字是： 今天有烤鸭吃，哈哈！")
//	//
//	//str := "今天有烤鸭吃，哈哈！ ! abc"
//	//str1 := Transer(str)
//	//fmt.Printf("中文变为unicode码， 略略略：“ 今天有烤鸭吃，哈哈！”变： %s", str1)
//	//fmt.Println()
//	//str2 := "小是正则表达式要匹配的字符串包含中英文vsdfvsva猪cas飞天demicgwegr//lwe;rgc了"
//	//fmt.Println(str2)
//	//fmt.Println("请匹配：字符串是否以“小”开头中间按顺序有“d”，有“猪”，以“了“结尾”")
//	//
//	//fmt.Println(strconv.QuoteToASCII("小")) //转换为unicode
//	//fmt.Println(strconv.QuoteToASCII("猪"))
//	//fmt.Println(strconv.QuoteToASCII("了")) //先答应出来查看对应转码，再在匹配规则中添加
//	////regex := `^[\u5c0f].*?d.*[\\u732a].*[\u4e86]$`
//	//regex := `[\\u5c0f]` + `.*[\\u732a].*d.*[\\u4e86]`
//	////regex := strconv.QuoteToASCII("小")
//	//fmt.Printf("匹配规则如下:\n%s", regex)
//	//s1 := regexp.MustCompile(regex)
//	//fmt.Println()
//	//fmt.Println(s1.MatchString(str2))
//	//fmt.Println(s1.MatchString(str1))
//	//a := strconv.QuoteToASCII("<fieldset style=\"line-height:150%;font-size:12px;\"><legend>&nbsp;错误，请联系管理员&nbsp;</legend><b>文件：</b>/general/file_folder/swfupload_new.php</fieldset>")
//	//fmt.Println(a)
//	fmt.Println(url.QueryEscape("id=1 union select 1,2,sys.fn_sqlvarbasetostr(HashBytes('MD5','abc')),db_name(1),5,6,7"))
//}

func main() {
	test1111()
}
func test2() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "53756725@qq.com"}
	fmt.Printf("m1: %v\n", m1) //	map是无序的
	m2 := make(map[string]string)
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["email"] = "53756725@qq.com"
	fmt.Printf("m2: %v\n", m2)
	fmt.Println(m2)
}
func Tongda() {
	fmt.Println("tongda")
}
func wanhu() {
	fmt.Println("wanhu")
}
func yongyou() {
	fmt.Println("yongyou")
}
func fanwei() {
	fmt.Println("fanwei")
}
func zhiyuan() {
	fmt.Println("zhiyuan")
}
func lanling() {
	fmt.Println("lanling")
}
func fanruan() {
	fmt.Println("fanruan")
}

var myMap = map[string]interface{}{
	"通达": Tongda,
	"泛微": fanwei,
}
var myMap1 = map[string][]string{
	"通达OA": {"通达OA SQL注入", "title=\"tongda\""},
	"泛微OA": {"泛微OA SQL注入", "title=\"fanwei\""},
}

func test1111() {
	for k := range myMap {
		fmt.Println(k)
	}
	for k, v := range myMap1 {
		fmt.Printf("%v:%v\n", k, v)
	}
}
