package util_test

import (
	"fmt"

	"github.com/wxw9868/util/pagination"
)

// ExampleNewPaginator 分页示例
func ExampleNewPaginator() {
	pages := pagination.NewPaginator(200, 10)
	pages.SetCurrentPage(10)

	fmt.Println(pages.TotalPage())
	fmt.Println(pages.PrePage())
	fmt.Println(pages.CurrentPage())
	fmt.Println(pages.NextPage())
	fmt.Println(pages.Pages())
	fmt.Println(pages)
}
