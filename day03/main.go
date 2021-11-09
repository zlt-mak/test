package main

////九九乘法表
//func main()  {
//	for i := 1; i < 10; i++ {
//		for j :=1; j <= i; j++ {
//			fmt.Printf("%d*%d=%d\t",j,i,j*i)
//		}
//		fmt.Println()
//	}
//
//}


//const (
//	a       = iota //a=0
//	b       = "B"
//	c       = iota             //c=2
//	d, e, f = iota, iota, iota //d=3,e=3,f=3
//	g       = iota             //g = 4
//)

var arr [10]int  // 声明了一个int类型的数组

func main()  {
	////当i=5时跳出for循环
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Println("over")
	//
	////当i=5时，跳过此次for循环（不执行for循环内部的打印语句）继续下一次循环
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		continue //继续下一次循环
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Println("over")

	//switch
	//switch n := 3; n {
	//case 1:
	//	fmt.Println("周一")
	//case 2:
	//	fmt.Println("周二")
	//case 3:
	//	fmt.Println("周三")
	//case 4:
	//	fmt.Println("周四")
	//case 5:
	//	fmt.Println("周五")
	//default:
	//	fmt.Println("无效")
	//}
	//
	//switch n :=5; n {
	//case 1,3,5,7,9:
	//	fmt.Println("奇数")
	//case 2,4,6,8,10:
	//	fmt.Println("偶数")
	//default:
	//	fmt.Println(n)
	//}

	//age := 22
	//if age >18 && age < 60 {
	//	fmt.Println("苦逼上班")
	//}else {
	//	fmt.Println("玩耍")
	//}
	//
	//if age <18 || age > 60 {
	//	fmt.Println("玩")
	//}else {
	//	fmt.Println("上班")
	//}
	//s := "hello"
	//s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	//fmt.Printf("%s\n", s)
	//fmt.Println(a, b, c, d, e, f, g)
	//arr[0] = 42      // 数组下标是从0开始的
	//arr[1] = 13
	//fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	//fmt.Printf("The last element is %d\n", arr[1]) //返回未赋值的最后一个元素，默认返回0


}