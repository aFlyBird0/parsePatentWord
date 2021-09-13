/*
组装前面获取的段落，加入层级属性，并存入数据库
 */
package build

import (
	"parsePatentWord/read"
	"parsePatentWord/save"
)

//回溯树
type trace struct {
	lastLevel int
	stack
}

type stack struct {
	cateIds []int "当前的层次最深的分类，在数据库中的分类号"
	top int
}

func Build()  {
	//docName := "\\static\\《专利审查指南》(2020年2月1日实施版).docx"
	docName := "\\static\\mini.docx"
	paras := read.Get(docName)
	// 创建追踪树，设定目前 lvl 为 -1,即树根
	t := &trace{-1, stack{make([]int, 0, 10), -1}}
	//获得整本书的目录
	rootId := save.SaveCate("专利审查指南", -1)
	t.stack.append(rootId)
	for _, para := range paras {
		curLvl := para.Lvl()
		curContent := para.Content()
		// 当前是文本段落
		if curLvl == 0 {
			save.SaveText(curContent, t.CateId())
		} else {
			//当前是分类
			// 层级变深
			if curLvl > t.lastLevel {
				id := save.SaveCate(curContent, t.CateId())
				t.stack.append(id)
				t.lastLevel++
			} else if curLvl == t.lastLevel {
				// 相同层级
				// 层级相同的话，说明到了个新的同层目录
				// 也就是说上面那个同级目录没用了
				t.stack.pop()
				id := save.SaveCate(curContent, t.CateId())
				t.stack.append(id)
			} else {
				// 层级变浅，即遇到更大的目录
				// 与上面类似， pop 两次即可
				// 第一次 pop 同级，第二次 pop 父级
				// 可以理解为, 每次 pop 一次， trace 的等级就加1（变浅1）
				// 最终要 pop 到 栈顶的值大于 现在
				t.stack.pop()
				t.stack.pop()
				id := save.SaveCate(curContent, t.CateId())
				t.stack.append(id)
				t.lastLevel--
				if t.lastLevel == 0 {
					t.lastLevel = -1
				}
			}
		}
	}
}
func (s *stack)append(id int)  {
	s.cateIds = append(s.cateIds, id)
	s.top++
}

func (s *stack)pop()  {
	s.cateIds = s.cateIds[:s.top]
	s.top--
}

func (s *stack)CateId()(cateId int)  {
	return s.cateIds[s.top]
}
