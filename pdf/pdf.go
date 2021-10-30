package pdf

import (
	"fmt"
	"os"

	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
)

const (
	TABLE_IG = "IPAexG"
	TABLE_MD = "MPBOLD"
	TABLE_MY = "微软雅黑"

	FONT_MY = "微软雅黑"
)

var (
	HeadFont = core.Font{Family: FONT_MY, Size: 12}
	TextFont = core.Font{Family: FONT_MY, Size: 10}
)

type PDF struct {
	FilePath string
	Title    string
	Fs       []Content
}

type Content func(report *core.Report)

func New(title string, fs []Content, filepath string) *PDF {
	return &PDF{
		FilePath: filepath,
		Fs:       fs,
		Title:    title,
	}
}

func (pdf *PDF) Run() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "/ttf/"

	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: TABLE_IG,
		FileName: path + "ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: TABLE_MD,
		FileName: path + "mplus-1p-bold.ttf",
	}
	font3 := core.FontMap{
		FontName: TABLE_MY,
		FileName: path + "microsoft.ttf",
	}

	r.SetFonts([]*core.FontMap{&font1, &font2, &font3})
	r.SetPage("A4", "P")
	r.FisrtPageNeedHeader = true
	r.FisrtPageNeedFooter = true

	r.RegisterExecutor(pdf.Executor, core.Detail)
	r.RegisterExecutor(pdf.ComplexReportFooterExecutor, core.Footer)
	r.RegisterExecutor(pdf.ComplexReportHeaderExecutor, core.Header)

	r.Execute(pdf.FilePath)
}

func (pdf *PDF) Executor(report *core.Report) {
	for _, f := range pdf.Fs {
		f(report)
	}
}

func (pdf *PDF) ComplexReportFooterExecutor(report *core.Report) {
	content := fmt.Sprintf("第 %v / {#TotalPage#} 页", report.GetCurrentPageNo())
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(TextFont)
	footer.SetFontColor("60, 179, 113")
	_ = footer.HorizontalCentered().SetContent(content).GenerateAtomicCell()
}
func (pdf *PDF) ComplexReportHeaderExecutor(report *core.Report) {
	footer := gopdf.NewSpan(10, 0, report)
	footer.SetFont(TextFont)
	footer.SetFontColor("255,0,0")
	footer.SetBorder(core.Scope{Top: 10})
	_ = footer.HorizontalCentered().SetContent(pdf.Title).GenerateAtomicCell()
}
