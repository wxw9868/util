package page

import (
	"fmt"
	"testing"
)

func TestNewPage(t *testing.T) {
	pages := NewPage(200, 10)
	pages.SetCurrentPage(10)

	fmt.Println(pages.TotalPage())
	fmt.Println(pages.PrePage())
	fmt.Println(pages.CurrentPage())
	fmt.Println(pages.NextPage())
	fmt.Println(pages.Pages())
	fmt.Println(pages)
}
