package parse

import (
	"baliance.com/gooxml/document"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
	"parsePatentWord/util"
	"strings"
)

func TestGooxmlParse() {
	filePath := util.GetRunPath()
	doc, err := document.Open(filePath + "\\static\\《专利审查指南》(2020年2月1日实施版).docx")
	//doc, err := document.Open(filePath + "\\static\\审查指南修改.docx")
	//doc, err := document.Open(filePath + "\\static\\mini.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	for _, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		var outlineLvl int64
	    if outlineLvlStruct := para.Properties().X().OutlineLvl; outlineLvlStruct!=nil{
			outlineLvl = outlineLvlStruct.ValAttr
		}else {
			outlineLvl = 0
		}
		fmt.Printf("大纲：%d\n", outlineLvl)
		//fmt.Printf("格式: %s\n", para.Style())
		//if style := para.Properties().Style(); style != "" {
		//	fmt.Println("-----------第", i, "段-------------")
		//	fmt.Printf("属性:%s\n", style)
		//} else {
		//	continue
		//}
		var paraStr string = ""
		for _, run := range para.Runs() {
			paraStr += run.Text()
			//fmt.Print("\t-----------第 j, "格式片段-------------")
			//fmt.Print("第", j, "格式片段:")
			//fmt.Println(run.Text())
			//para.SetStyle()
		}
		fmt.Println(paraStr)
	}
}

func TestNewDocGooxml()  {
	filePath := util.GetRunPath()
	doc := document.New()
	p := doc.AddParagraph()
	r := p.AddRun()
	p.SetStyle("Heading1")
	r.AddText("这是我加的标题一")
	doc.SaveToFile(filePath + "\\static\\testNewDoc.doc")
}
type outlineLevel string

const (
	h1 outlineLevel = "\n1"
	h2 outlineLevel = "2"
	h3 outlineLevel = "3"
	h4 outlineLevel = "4"
	h5 outlineLevel = "5"
	h6 outlineLevel = "6"
)

type Doc struct {
	*html.Node
}

func TestHtmlQuery()  {
	filePath := util.GetRunPath()
	filename := filePath + "\\static\\《专利审查指南》(2020年2月1日实施版).htm"
	//filename := filePath + "\\static\\mini.htm"
	doc, err := htmlquery.LoadDoc(filename)
	if err != nil {
		fmt.Println(err)
	}
	//getAllH1(doc)
	//getAllH2(doc)
	myDoc := Doc{doc}
	//myDoc.getContents(h1)
	//myDoc.getContents(h2)
	myDoc.getContents(h3)
	//myDoc.getContents(h4)
	//myDoc.getContents(h5)
	//myDoc.getContents(h6)
}

func getAllH1(doc *html.Node)  {
	firstCates := htmlquery.Find(doc, "//p[contains(@style,'mso-outline-level:\n1')]")
	for _, firstCate := range firstCates{
		//fmt.Printf("%d:%v\n", index, item)
		firstCateList := htmlquery.Find(firstCate, "//text()")
		var firstCateStr string = ""
		for _, item := range firstCateList{
			//fmt.Printf("%d:%v", indexJ, item)
			firstCateStr += item.Data
		}
		firstCateStr = strings.TrimSpace(firstCateStr)
		if  firstCateStr!= ""{
			fmt.Println(firstCateStr)
		}

		//firstCateString := strings.Join(, "")
	}
	fmt.Println(len(firstCates))
}


func (doc Doc)getContents(level outlineLevel )  {
	expr := "//p[contains(@style,'mso-outline-level:"+string(level)+"')]"
	cates := htmlquery.Find(doc.Node, expr)
	for _, cate := range cates{
		//fmt.Printf("%d:%v\n", index, item)
		cateList := htmlquery.Find(cate, "//span/text()")
		//cateList := htmlquery.Find(cate, "//*[not(table)]/span[not(contains(@style, 'left:0pt'))]//text()")
		//cateList := htmlquery.Find(cate, "//*[not(table)]/span//text()")
		//cateList := htmlquery.Find(cate, "//span[contains(@style, 'font-family:黑体')]/text()")

		var cateStr string
		for _, item := range cateList{
			//fmt.Printf("%d:%v", indexJ, item)
			cateStr += item.Data
		}
		cateStr = strings.TrimSpace(cateStr)
		if  cateStr!= ""{
			fmt.Println("爱飞的鸟 " + strings.ReplaceAll(cateStr, " ", ""))
		}

		//firstCateString := strings.Join(, "")
	}
	fmt.Printf("共%d条",len(cates))
}


