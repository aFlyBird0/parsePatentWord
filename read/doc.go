/*
读取 docx 文件，合并相同层级段落
 */
package read

type Doc struct {
	//读取出的文档类
	paras Paras
	filename string	//相对于主目录的存储路径
}

type Paras []*Para

type Para struct {
	// 封装了大纲和内容的类
	outlineLvl int64
	content string
}

func newDoc(filename string) (doc *Doc) {
	return &Doc{filename: filename}
}
