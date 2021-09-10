package parse

import (
	"testing"
)

func TestEtree(t *testing.T){
	//doc := etree.NewDocument()
	//filePath := util.GetRunPath()
	//// 因为 html 有些节点不是闭合的，所以用 etree 读取会失败
	//if err := doc.ReadFromFile(filePath + "\\static\\mini.htm"); err != nil {
	//	panic(err)
	//}
	//fmt.Println(doc)
}

//func TestHtmlQuery(t *testing.T)  {
//	filePath := util.GetRunPath()
//	filename := filePath + "\\static\\《专利审查指南》(2020年2月1日实施版).htm"
//	doc, err := htmlquery.LoadDoc(filename)
//	if err != nil {
//		fmt.Println(err)
//	}
//	firstCates := htmlquery.Find(doc, "//p[contains(@style,'mso-outline-level:\n1')]")
//	for _, firstCate := range firstCates{
//		//fmt.Printf("%d:%v\n", index, item)
//		firstCateList := htmlquery.Find(firstCate, "//text()")
//		for indexJ, item := range firstCateList{
//			fmt.Printf("%d:%v", indexJ, item)
//		}
//
//		//firstCateString := strings.Join(, "")
//	}
//	fmt.Println(len(firstCates))
//
//}
