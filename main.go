package main

import (
	"fmt"
	"parsePatentWord/build"
)

func main() {
	fmt.Println("解析开始")
	//read.TestGooxmlParse()
	//read.TestEtree()
	//read.TestHtmlQuery()
	//read.TestNewDocGooxml()
	docName := "\\static\\《专利审查指南》(2020年2月1日实施版).docx"
	//docName := "\\static\\mini.docx"
	build.Build(docName)
	fmt.Println("保存完成")
}
