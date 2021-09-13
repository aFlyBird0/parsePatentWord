package read

import (
	"fmt"
	"strings"
)

func (paras Paras) Combine() (parasReturn Paras) {
	//合并相同层级段
	parasReturn = make(Paras, 0, len(paras))
	//parasReturn = make(Paras, len(paras))
	var accContent strings.Builder
	lastLvl := -1
	for _, para := range paras {
		curLvl := para.outlineLvl
		curContent := para.content
		// 如果上一个 outlineLevel 和当前的不同，就把之前累积的字符存进去
		if lastLvl != curLvl {
			parasReturn = append(parasReturn, &Para{lastLvl, strings.TrimSpace(accContent.String())})
			accContent.Reset()
		}
		// 无论是否相同，都累积字符，区别在于一个要清空一个不清空
		accContent.WriteString(curContent)
		accContent.WriteString("\n")
		lastLvl = curLvl
		//fmt.Println(para)
	}
	fmt.Printf("目录+段落一共 %d 条", len(parasReturn)-1)
	return parasReturn[1:]
}
