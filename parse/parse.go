package parse

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
	"parsePatentWord/util"
)

type Para struct {
	// 封装了大纲和内容的类
	outlineLvl int64
	content string
}

func (p Para)String() string {
	return fmt.Sprintf("%v : %s", p.outlineLvl, p.content)
}
func getParaOutlineAndContent() []Para {
	filePath := util.GetRunPath()
	doc, err := document.Open(filePath + "\\static\\《专利审查指南》(2020年2月1日实施版).docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	length := len(doc.Paragraphs())
	paras := make([]Para, length, length)
	for index, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		var outlineLvl int64
		if outlineLvlStruct := para.Properties().X().OutlineLvl; outlineLvlStruct!=nil{
			outlineLvl = outlineLvlStruct.ValAttr
		}else {
			outlineLvl = 0
		}
		//fmt.Printf("大纲：%d\n", outlineLvl)
		var paraStr string = ""
		for _, run := range para.Runs() {
			paraStr += run.Text()
		}
		//fmt.Println(paraStr)
		if paraStr != "" {
			paras[index] = Para{outlineLvl, paraStr}
		}

	}
	return paras
}

func TestParse()  {
	paras := getParaOutlineAndContent()
	for _, para := range paras{
		fmt.Println(para)
	}
}
