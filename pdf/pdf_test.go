package pdf

import (
	"testing"
)

func TestNew(t *testing.T) {
	//titles := []string{"日期", "发货地", "收货地", "运单号", "物流手机号"}
	//cols := len(titles)
	//lineSpace := 1.0
	//lineHeight := 18.0
	//f := func(table *gopdf.Table, report *core.Report) {
	//	f1 := core.Font{Family: TABLE_MY, Size: 8}
	//	border := core.NewScope(4.0, 4.0, 0, 0)
	//
	//	for i := 0; i < 1-1; i++ {
	//		cells := make([]*gopdf.TableCell, cols)
	//		for j := 0; j < cols; j++ {
	//			cells[j] = table.NewCell()
	//		}
	//
	//		for j := 0; j < cols; j++ {
	//			str := `有限公司送达行政处罚决定书`
	//			s := fmt.Sprintf("%v-%v", i+2, str)
	//			w := table.GetColWidth(0, j)
	//			e := gopdf.NewTextCell(w, lineHeight, lineSpace, report)
	//			e.SetFont(f1)
	//			if i%2 == 0 {
	//				e.SetBackColor("255,192,203")
	//			}
	//			e.SetBorder(border)
	//			e.SetContent(s)
	//			cells[j].SetElement(e)
	//		}
	//	}
	//}

	//pdf := New("223.pdf", cols, 1, 400, titles, f，nil，nil)
	//pdf.Run()
}
