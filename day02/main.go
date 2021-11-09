package main

import "fmt"

func main()  {
	//path:="\"D:\\Go_code\\src\\code.oldboyedu.com\\studygo\\day02\""
	//fmt.Println(path)
	////多行字符串
	//s2 :=`
	//	张里涛
	//	张瑞
	//`
	//fmt.Println(s2)
	//
	//s3 :=`D:\Go_code\src\code.oldboyedu.com\studygo\day02`
	//fmt.Println(s3)
	//
	////字符串拼接
	//name := "张里"
	//word := "涛"
	//ss := name + word
	//
	//fmt.Println(ss)
	//
	//ss1 := fmt.Sprintf("%s%s",name,word)
	//fmt.Println(ss1)
	//
	////分割
	//ret := strings.Split(s3,"\\")
	//fmt.Println(ret)
	//
	////判断是否包含
	//fmt.Println(strings.Contains(ss,"张里涛1"))
	//
	//age := 19
	//if age > 35 {
	//	fmt.Println("成功")
	//}else if age > 18 {
	//	fmt.Println("青年")
	//} else {
	//	fmt.Println("失败")
	//}
	//
	////基本格式
	//for i :=0; i < 10; i++{
	//	fmt.Println(i)
	//}

	//for range循环
	s := "Hello沙河"
	for i, v := range s{
		//fmt.Printf("%d %c\n", i, v)
		fmt.Println(i,v)
	}
}
