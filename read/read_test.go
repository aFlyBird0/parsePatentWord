package read

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("test read docx", func(t *testing.T) {
		paras := Read("\\static\\《专利审查指南》(2020年2月1日实施版).docx")
		for _, para := range paras {
			fmt.Println(para)
		}
	})
}

func TestCombine(t *testing.T) {
	t.Run("test combine paras", func(t *testing.T) {
		paras := Read("\\static\\《专利审查指南》(2020年2月1日实施版).docx")
		parasCombine := paras.Combine()
		fmt.Println(parasCombine)
	})
}
