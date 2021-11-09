package main

import "fmt"

////声明变量
//var(
//	name string
//	age  int
//	isok bool
//)

////程序的入口函数
//func main()  {
//	name="zhang"
//	age=21
//	isok=true
//	fmt.Print(isok)
//	fmt.Printf("name:%s\n",name)
//	fmt.Println(age)
//}

//定义常量
const (
	statusOk = 200
	notFount = 404
)

const (
	n1 = 100
	n2
	n3
)

func main()  {
	fmt.Println("n1:",n1)
	fmt.Println("n2:",n2)
	fmt.Println("n3:",n3)
}