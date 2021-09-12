package page

import (
	"math"
)

type Page struct {
	totalCount  int   //总条数
	totalPage   int   //总页数
	prePage     int   //上一页
	currentPage int   //当前页
	nextPage    int   //下一页
	pageSize    int   //每页数量
	pageRange   []int //页数列表
}

//初始化结构体
func NewPage(totalCount, pageSize int) *Page {
	if pageSize <= 0 {
		pageSize = 10
	}
	return &Page{
		totalCount: totalCount,
		pageSize:   pageSize,
	}
}

//设置当前页
func (p *Page) SetCurrentPage(currentPage int) {
	if currentPage > 0 {
		p.currentPage = currentPage
	}
	if currentPage <= 0 {
		p.currentPage = 1
	}
	if currentPage > p.TotalPage() {
		p.currentPage = p.TotalPage()
	}
}

//总页数
func (p *Page) TotalPage() int {
	if p.totalPage > 0 {
		return p.totalPage
	}
	p.totalPage = int(math.Ceil(float64(p.totalCount) / float64(p.pageSize)))
	return p.totalPage
}

//首页
func (p *Page) FirstPage() int {
	return 1
}

//尾页
func (p *Page) LastPage() int {
	if p.totalPage > 0 {
		return p.totalPage
	}
	return p.TotalPage()
}

//上一页
func (p *Page) PrePage() int {
	if p.prePage > 0 {
		return p.prePage
	}
	p.prePage = int(math.Max(float64(1), float64(p.CurrentPage()-1)))
	return p.prePage
}

//当前页
func (p *Page) CurrentPage() int {
	if p.currentPage > 0 {
		return p.currentPage
	}
	return 1

}

//下一页
func (p *Page) NextPage() int {
	if p.nextPage > 0 {
		return p.nextPage
	}
	p.nextPage = int(math.Min(float64(p.TotalPage()), float64(p.CurrentPage()+1)))
	return p.nextPage
}

//如果当前页有上一页，HasPrev将返回true。
func (p *Page) HasPrev() bool {
	return p.CurrentPage() > 1
}

//如果当前页有下一页，HasNext将返回true。
func (p *Page) HasNext() bool {
	return p.CurrentPage() < p.TotalPage()
}

//如果给定的页索引指向当前页，IsActive返回true。
func (p *Page) IsActive(page int) bool {
	return p.CurrentPage() == page
}

//Offset返回当前偏移量。
func (p *Page) Offset() int {
	return (p.CurrentPage() - 1) * p.pageSize
}

//如果有多个页面，HasPages返回true。
func (p *Page) HasPages() bool {
	return p.TotalPage() > 1
}

//Pages返回所有页面的列表。
func (p *Page) Pages() []int {
	if p.pageRange == nil && p.totalCount > 0 {
		var pages []int
		pageNums := p.TotalPage()
		page := p.CurrentPage()
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}
