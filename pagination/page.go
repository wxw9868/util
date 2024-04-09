package pagination

import (
	"math"
)

type Paginator struct {
	totalCount  int   //总条数
	totalPage   int   //总页数
	prePage     int   //上一页
	currentPage int   //当前页
	nextPage    int   //下一页
	pageSize    int   //每页数量
	pageRange   []int //页数列表
}

// NewPaginator 初始化结构体
func NewPaginator(totalCount, pageSize int) *Paginator {
	if pageSize <= 0 {
		pageSize = 10
	}
	return &Paginator{
		totalCount: totalCount,
		pageSize:   pageSize,
	}
}

// SetCurrentPage 设置当前页
func (p *Paginator) SetCurrentPage(currentPage int) {
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

// TotalPage 总页数
func (p *Paginator) TotalCount() int {
	if p.totalCount > 0 {
		return p.totalCount
	}
	return p.totalCount
}

// TotalPage 总页数
func (p *Paginator) TotalPage() int {
	if p.totalPage > 0 {
		return p.totalPage
	}
	p.totalPage = int(math.Ceil(float64(p.totalCount) / float64(p.pageSize)))
	return p.totalPage
}

// FirstPage 首页
func (p *Paginator) FirstPage() int {
	return 1
}

// LastPage 尾页
func (p *Paginator) LastPage() int {
	if p.totalPage > 0 {
		return p.totalPage
	}
	return p.TotalPage()
}

// PrePage 上一页
func (p *Paginator) PrePage() int {
	if p.prePage > 0 {
		return p.prePage
	}
	p.prePage = int(math.Max(float64(1), float64(p.CurrentPage()-1)))
	return p.prePage
}

// CurrentPage 当前页
func (p *Paginator) CurrentPage() int {
	if p.currentPage > 0 {
		return p.currentPage
	}
	return 1

}

// NextPage 下一页
func (p *Paginator) NextPage() int {
	if p.nextPage > 0 {
		return p.nextPage
	}
	p.nextPage = int(math.Min(float64(p.TotalPage()), float64(p.CurrentPage()+1)))
	return p.nextPage
}

// HasPrev 如果当前页有上一页，HasPrev将返回true。
func (p *Paginator) HasPrev() bool {
	return p.CurrentPage() > 1
}

// HasNext 如果当前页有下一页，HasNext将返回true。
func (p *Paginator) HasNext() bool {
	return p.CurrentPage() < p.TotalPage()
}

// IsActive 如果给定的页索引指向当前页，IsActive返回true。
func (p *Paginator) IsActive(page int) bool {
	return p.CurrentPage() == page
}

// Offset Offset返回当前偏移量。
func (p *Paginator) Offset() int {
	return (p.CurrentPage() - 1) * p.pageSize
}

// HasPages 如果有多个页面，HasPages返回true。
func (p *Paginator) HasPages() bool {
	return p.TotalPage() > 1
}

// Pages Pages返回所有页面的列表。
func (p *Paginator) Pages() []int {
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
