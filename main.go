package main

import (
	"fmt"
	"parsePatentWord/build"
)

func main()  {
	fmt.Println("解析开始")
	//read.TestGooxmlParse()
	//read.TestEtree()
	//read.TestHtmlQuery()
	//read.TestNewDocGooxml()
	build.Build()
	fmt.Println("保存完成")
}
