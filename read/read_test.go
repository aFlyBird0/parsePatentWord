package read

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T)  {
	t.Run("test read docx", func(t *testing.T) {
		papas := Read()
		for _, para := range papas{
			fmt.Println(para)
		}
	})
}
