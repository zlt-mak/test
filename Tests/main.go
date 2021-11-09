package main

import (
	"flag"
	"fmt"
)

func main()  {
	var t int
	flag.IntVar(&t,"t",1,"1唐诗2宋词")
	flag.Parse()
	if t == 1 {
		fmt.Println("唐诗")
	}else if t == 2{
		fmt.Println("宋词")
	}
}
