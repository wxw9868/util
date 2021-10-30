package pdf

import (
	"fmt"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
)

var lineSpace, lineHeight = 1.0, 13.0

func titleFunc(report *core.Report, content string) {
	report.SetMargin(0, 20)
	baseInfoDiv := gopdf.NewDivWithWidth(600, lineHeight, lineSpace, report)
	baseInfoDiv.SetFont(HeadFont)
	_ = baseInfoDiv.SetContent(content).GenerateAtomicCell()
	report.SetMargin(0, 10)
}

func tableFunc(report *core.Report, content string, titles []string, cols, rows int) *gopdf.Table {
	titleFunc(report, content)

	table := gopdf.NewTable(cols, rows, 600, lineHeight, report)
	table.SetMargin(core.Scope{})
	font := core.Font{Family: TABLE_MY, Size: 8, Style: ""}
	border := core.NewScope(4.0, 4.0, 4.0, 3.0)
	for _, title := range titles {
		r := table.NewCellByRange(1, 1)
		e := gopdf.NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report)
		e.SetFont(font).SetBorder(border).HorizontalCentered().SetContent(title)
		r.SetElement(e)
	}
	return table
}

func TestNew(t *testing.T) {
	f1 := func(report *core.Report) {
		report.SetMargin(0, -20)
		baseInfoDiv := gopdf.NewDivWithWidth(500, lineHeight, lineSpace, report)
		baseInfoDiv.SetFont(HeadFont)
		_ = baseInfoDiv.SetContent(fmt.Sprintf("订单编号: %s", "202110290001")).GenerateAtomicCell()

		report.SetMargin(0, 20)
		baseInfoDiv = gopdf.NewDivWithWidth(500, lineHeight, lineSpace, report)
		baseInfoDiv.SetFont(HeadFont)
		_ = baseInfoDiv.SetContent("客户信息").GenerateAtomicCell()

		report.SetMargin(0, 10)

		baseInfo := gopdf.NewDivWithWidth(500, lineHeight, lineSpace, report)
		baseInfo.SetMarign(core.Scope{Left: 0, Top: 2.0})
		_ = baseInfo.SetFont(TextFont).SetContent(fmt.Sprintf("客户: %s", "萧炎")).GenerateAtomicCell()
		_ = baseInfo.Copy(fmt.Sprintf("手机号: %s", "18266668888")).GenerateAtomicCell()
	}

	res := []map[string]interface{}{
		{
			"ship_date":        int64(0),
			"ship_address":     "",
			"delivery_address": "",
			"tracking_number":  "",
			"mobile":           "18266668888",
		},
	}
	f2 := func(report *core.Report) {
		content := "物流信息"
		titles := []string{"日期", "发货地", "收货地", "运单号", "物流手机号"}
		cols, rows := len(titles), len(res)+1
		table := tableFunc(report, content, titles, cols, rows)

		font := core.Font{Family: TABLE_MY, Size: 8}
		border := core.NewScope(4.0, 4.0, 0, 0)
		var column []string
		for i := 0; i < rows-1; i++ {
			cells := make([]*gopdf.TableCell, cols)
			for j := 0; j < cols; j++ {
				cells[j] = table.NewCell()
			}

			shipDate := time.Unix(res[i]["ship_date"].(int64), 0).Format("2006-01-02")
			column = append(column, shipDate)
			column = append(column, res[i]["ship_address"].(string))
			column = append(column, res[i]["delivery_address"].(string))
			column = append(column, res[i]["tracking_number"].(string))
			column = append(column, res[i]["mobile"].(string))

			for j := 0; j < cols; j++ {
				w := table.GetColWidth(0, j)
				e := gopdf.NewTextCell(w, lineHeight, lineSpace, report)
				e.SetFont(font)
				if i%2 == 0 {
					e.SetBackColor("255,192,203")
				}
				e.SetBorder(border)
				e.SetContent(column[j])
				cells[j].SetElement(e)
			}
		}
		_ = table.GenerateAtomicCell()
	}

	fs := []Content{f1, f2}
	filePath := "./" + strconv.Itoa(int(time.Now().Unix())) + ".pdf"
	New("物流记录", fs, filePath).Run()
}

func TestNew2(t *testing.T) {
	f1 := func(report *core.Report) {
		report.SetMargin(0, -20)
		baseInfoDiv := gopdf.NewDivWithWidth(500, lineHeight, lineSpace, report)
		baseInfoDiv.SetFont(HeadFont)
		_ = baseInfoDiv.SetContent(fmt.Sprintf("订单编号: %s", "202110290001")).GenerateAtomicCell()

		report.SetMargin(0, 20)
		baseInfoDiv = gopdf.NewDivWithWidth(500, lineHeight, lineSpace, report)
		baseInfoDiv.SetFont(HeadFont)
		_ = baseInfoDiv.SetContent("客户信息").GenerateAtomicCell()

		report.SetMargin(0, 10)

		baseInfo := gopdf.NewDivWithWidth(500, lineHeight, lineSpace, report)
		baseInfo.SetMarign(core.Scope{Left: 0, Top: 2.0})
		_ = baseInfo.SetFont(TextFont).SetContent(fmt.Sprintf("客户: %s", "萧炎")).GenerateAtomicCell()
		_ = baseInfo.Copy(fmt.Sprintf("手机号: %s", "18266668888")).GenerateAtomicCell()
	}

	// 矿机信息
	minerInfo := []map[string]interface{}{
		{
			"brand_name":    "",
			"brand_model":   "",
			"hashrate":      float64(0),
			"hashrate_unit": "",
			"power":         int64(0),
			"miner_count":   uint32(0),
			"put_count":     int64(0),
		},
	}
	f2 := func(report *core.Report) {
		content := "矿机信息"
		titles := []string{"矿机品牌", "矿机型号", "额定算力", "额定功率（W）", "订单矿机数量", "入库矿机数量"}
		cols, rows := len(titles), len(minerInfo)+1
		table := tableFunc(report, content, titles, cols, rows)

		font := core.Font{Family: TABLE_MY, Size: 8}
		border := core.NewScope(4.0, 4.0, 0, 0)
		var column []string
		for i := 0; i < rows-1; i++ {
			cells := make([]*gopdf.TableCell, cols)
			for j := 0; j < cols; j++ {
				cells[j] = table.NewCell()
			}

			column = append(column, minerInfo[i]["brand_name"].(string))
			column = append(column, minerInfo[i]["brand_model"].(string))
			float32s2 := strconv.FormatFloat(minerInfo[i]["hashrate"].(float64), 'f', -1, 64)
			column = append(column, float32s2+minerInfo[i]["hashrate_unit"].(string))
			column = append(column, strconv.Itoa(int(minerInfo[i]["power"].(int64))))
			column = append(column, strconv.Itoa(int(minerInfo[i]["miner_count"].(uint32))))
			column = append(column, strconv.Itoa(int(minerInfo[i]["put_count"].(int64))))

			for j := 0; j < cols; j++ {
				e := gopdf.NewTextCell(table.GetColWidth(0, j), lineHeight, lineSpace, report)
				e.SetFont(font)
				if i%2 == 0 {
					e.SetBackColor("255,192,203")
				}
				e.SetBorder(border)
				e.SetContent(column[j])
				cells[j].SetElement(e)
			}
		}
		_ = table.GenerateAtomicCell()
	}

	// 电源信息
	powerSupplyInfo := []map[string]interface{}{
		{
			"model":     "",
			"put_count": int64(0),
		},
	}
	f3 := func(report *core.Report) {
		content := "电源信息"
		titles := []string{"电源型号", "入库电源数量"}
		cols, rows := len(titles), len(powerSupplyInfo)+1
		table := tableFunc(report, content, titles, cols, rows)

		font := core.Font{Family: TABLE_MY, Size: 8}
		border := core.NewScope(4.0, 4.0, 0, 0)
		var column []string
		for i := 0; i < rows-1; i++ {
			cells := make([]*gopdf.TableCell, cols)
			for j := 0; j < cols; j++ {
				cells[j] = table.NewCell()
			}

			column = append(column, powerSupplyInfo[i]["model"].(string))
			column = append(column, strconv.Itoa(int(powerSupplyInfo[i]["put_count"].(int64))))

			for j := 0; j < cols; j++ {
				w := table.GetColWidth(0, j)
				e := gopdf.NewTextCell(w, lineHeight, lineSpace, report)
				e.SetFont(font)
				if i%2 == 0 {
					e.SetBackColor("255,192,203")
				}
				e.SetBorder(border)
				e.SetContent(column[j])
				cells[j].SetElement(e)
			}
		}
		_ = table.GenerateAtomicCell()
	}

	// 出库确认单
	dir, _ := filepath.Abs("pictures")
	filename := fmt.Sprintf("%v/qrcode.png", dir)
	f4 := func(report *core.Report) {
		titleFunc(report, "出库确认单")

		im := gopdf.NewImageWithWidthAndHeight(filename, 100, 500, report)
		im.SetMargin(core.Scope{Left: 0, Top: 0})
		im.SetAutoBreak()
		_, _, _ = im.GenerateAtomicCell()
	}

	fs := []Content{f1, f2, f3, f4}
	filePath := "./" + strconv.Itoa(int(time.Now().Unix())) + ".pdf"
	New("出库单", fs, filePath).Run()
}
