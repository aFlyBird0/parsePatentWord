package read

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T)  {
	t.Run("test read docx", func(t *testing.T) {
		paras := Read()
		for _, para := range paras{
			fmt.Println(para)
		}
	})
}

func TestCombine(t *testing.T)  {
	t.Run("test combine paras", func(t *testing.T) {
		paras := Read()
		paras.Combine()
	})
}
