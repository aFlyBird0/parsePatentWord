package read

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
	"parsePatentWord/util"
)

type Doc struct {
	//读取出的文档类
	paras []*Para
	filename string	//相对于主目录的存储路径
}

type Para struct {
	// 封装了大纲和内容的类
	outlineLvl int64
	content string
}

func newDoc(filename string) (doc *Doc) {
	return &Doc{filename: filename}
}

func (doc *Doc)Get()(paras []*Para)  {
	return doc.paras
}


func (p Para)String() string {
	return fmt.Sprintf("%v : %s", p.outlineLvl, p.content)
}
func (myDoc *Doc)getParaOutlineAndContent() {
	filePath := util.GetRunPath()
	doc, err := document.Open(filePath + myDoc.filename)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	length := len(doc.Paragraphs())
	myDoc.paras = make([]*Para, 0, length)
	for _, para := range doc.Paragraphs() {
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
			myDoc.paras = append(myDoc.paras, &Para{outlineLvl, paraStr})
		}

	}
}

func Read()  (paras []*Para){
	doc := newDoc("\\static\\《专利审查指南》(2020年2月1日实施版).docx")
	doc.getParaOutlineAndContent()
	return doc.Get()
}