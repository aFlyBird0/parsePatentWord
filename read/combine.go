package read

import "fmt"

func (paras Paras)Combine()(parasReturn Paras)  {
	//合并相同层级段
	parasReturn = make(Paras, 0, len(paras))
	//parasReturn = make(Paras, len(paras))
	for _, para := range paras{
		fmt.Println(para)
	}
	return
}
