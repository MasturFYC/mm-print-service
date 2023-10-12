package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
	language "golang.org/x/text/language"
	message "golang.org/x/text/message"
)

func PrintNotaPdf(c echo.Context) error {

	data := CustomerOrder{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	var buf bytes.Buffer
	err := create_nota_pdf(&buf, &data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			HelloWorld{
				Message: fmt.Sprintf("Error. %v", err.Error()),
			})
	}
	c.Response().Writer.Header().Set("Content-Type", "application/pdf")
	c.Response().Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=nota-%d.pdf", data.ID))
	return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
}

func create_nota_pdf(writer io.Writer, data *CustomerOrder) (err error) {
	const (
		unit        = "mm"
		orientation = "P"
		saxmono     = "saxmono"
		//ocra        = "ocrb"
	)
	//var lh float64 = 16
	pgw := 95.0
	pgh := 140.0
	var mt float64 = 2.0
	var ml float64 = 2.0
	var mr float64 = 2.0

	dir, _ := os.Getwd()

	p := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        unit,
		Size:           gofpdf.SizeType{Wd: pgw, Ht: pgh},
		OrientationStr: orientation,
		FontDirStr:     filepath.Join(dir, "fonts"),
	})

	//log.Println(filepath.Join(dir, "fonts"))

	p.AddUTF8Font(saxmono, "", "saxmono.ttf")
	//p.AddUTF8Font(ocra, "B", "consolab.ttf")
	p.SetMargins(ml, mt, mr)
	p.SetAutoPageBreak(true, 3.0)

	//p.SetFillColor(255, 255, 255)
	p.SetTextColor(0, 0, 0)

	y := mt
	x := ml

	lnHeight := 3.5

	fileLogo := filepath.Join(dir, "tokomm.png")

	half := 95.0 / 3.0

	p.AddPage()
	p.SetFont(saxmono, "", 8)
	p.Image(fileLogo, x, y, half, 0, false, "", 0, "")

	col0 := 20.0
	x += half + 3.0
	p.SetXY(x, y+5.0)

	p.CellFormat(col0, lnHeight, "ORDER ID:", "", 0, "LT", false, 0, "")
	p.CellFormat(0, lnHeight, fmt.Sprintf("%d", data.ID), "", 1, "LT", false, 0, "")

	p.SetXY(x, p.GetY())
	p.CellFormat(col0, lnHeight, "SALES:", "", 0, "LT", false, 0, "")
	p.CellFormat(0, lnHeight, data.SalesName, "", 1, "LT", false, 0, "")

	p.SetXY(x, p.GetY())
	p.CellFormat(col0, lnHeight, "PELANGGAN:", "", 0, "LT", false, 0, "")
	p.CellFormat(0, lnHeight, data.CustomerName, "", 1, "LT", false, 0, "")

	p.SetXY(x, p.GetY())
	p.CellFormat(col0, lnHeight, "ALAMAT:", "", 0, "LT", false, 0, "")
	p.CellFormat(0, lnHeight, data.Address, "", 1, "LT", false, 0, "")

	x = ml
	p.SetXY(x, p.GetY()+2.0)
	p.Line(x, p.GetY(), pgw-x, p.GetY())
	//	p.Line(x, pgh-mt, pgw-x, pgh-mt)
	p.SetXY(x, p.GetY()+1.0)

	col1 := 35.0
	col2 := 8.0
	col3 := 8.0
	col4 := 19.0
	col5 := 21.0

	p.CellFormat(col1, lnHeight, "NAMA BARANG", "", 0, "LT", false, 0, "")
	p.CellFormat(col2, lnHeight, "QTY", "", 0, "RT", false, 0, "")
	p.CellFormat(col3, lnHeight, "UNIT", "", 0, "LT", false, 0, "")
	p.CellFormat(col4, lnHeight, "HARGA", "", 0, "RT", false, 0, "")
	p.CellFormat(col5, lnHeight, "SUBTOTAL", "", 1, "RT", false, 0, "")
	p.Line(x, p.GetY(), pgw-x, p.GetY())
	p.SetXY(x, p.GetY()+1.0)

	m := message.NewPrinter(language.Indonesian)

	for _, r := range data.Details {

		name := r.Name
		if len(r.VariantName) > 1 {
			name += ", " + r.VariantName
		}

		if len(name) > 20 {
			name = name[0:20]
		}

		p.SetXY(x, p.GetY())
		p.CellFormat(col1, lnHeight, name, "", 0, "LT", false, 0, "")
		p.CellFormat(col2, lnHeight, fmt.Sprintf("%v", r.Qty), "", 0, "RT", false, 0, "")
		p.CellFormat(col3, lnHeight, r.Unit, "", 0, "LT", false, 0, "")
		p.CellFormat(col4, lnHeight, m.Sprintf("%0.f", r.Pot), "", 0, "RT", false, 0, "")
		p.CellFormat(col5, lnHeight, m.Sprintf("%0.f", r.Subtotal), "", 1, "RT", false, 0, "")
	}

	p.Line(x, p.GetY(), pgw-x, p.GetY())
	p.SetXY(x, p.GetY()+2.0)
	p.CellFormat(col1+col2+col3, lnHeight, "TOTAL", "", 0, "RT", false, 0, "")
	p.SetFont(saxmono, "", 10)
	p.CellFormat(col4+col5, lnHeight, m.Sprintf("%0.f", data.Total), "", 1, "RB", false, 0, "")
	p.SetXY(x, p.GetY()+1.0)
	p.SetFont(saxmono, "", 8)
	p.CellFormat(col1+col2+col3, lnHeight, "BAYAR", "", 0, "RT", false, 0, "")
	p.CellFormat(col4+col5, lnHeight, m.Sprintf("%0.f", data.Payment), "", 1, "RT", false, 0, "")
	p.SetXY(x, p.GetY())
	p.CellFormat(col1+col2+col3, lnHeight, "SISA BAYAR", "", 0, "RT", false, 0, "")
	p.CellFormat(col4+col5, lnHeight, m.Sprintf("%0.f", data.RemainPayment), "", 1, "RT", false, 0, "")

	err = p.Output(writer)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	p.Close()
	//_ = pdf.OutputFileAndClose("hello.pdf")
	return nil
}
