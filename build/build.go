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
	top     int
}

func Build(docName string) {
	paras := read.Get(docName)
	// 创建追踪树，设定目前 lvl 为 -1,即树根
	t := &trace{0, stack{make([]int, 0, 10), -1}}
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
				// 层级增加的时候，每次一定是加 1
				t.lastLevel++
			} else if curLvl == t.lastLevel {
				// 相同层级
				// 层级相同的话，说明到了个新的同层目录
				// 也就是说上面那个同级目录没用了
				t.stack.pop()
				id := save.SaveCate(curContent, t.CateId())
				t.stack.append(id)
			} else {
				// 层级变浅，例如从 3 级目录变成 1 级目录
				// 层级减少的时候，不一定是减1，可能直接从 3 级目录变成 1 级了，这时候要减2
				// 最终要 pop 到 栈顶的值刚好比现在大 1
				lvlMinus := t.lastLevel - curLvl
				// 如果层级减1，pop 两次；减2，3次
				//之前相同的时候，可以理解为层级减0，pop 1次
				for i := 0; i < lvlMinus+1; i++ {
					t.stack.pop()
				}
				id := save.SaveCate(curContent, t.CateId())
				t.stack.append(id)
				t.lastLevel -= lvlMinus
			}
		}
	}
}
func (s *stack) append(id int) {
	s.cateIds = append(s.cateIds, id)
	s.top++
}

func (s *stack) pop() {
	s.cateIds = s.cateIds[:s.top]
	s.top--
}

func (s *stack) CateId() (cateId int) {
	return s.cateIds[s.top]
}
