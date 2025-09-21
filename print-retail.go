package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"strings"

	"github.com/alexbrainman/printer"
	"github.com/labstack/echo/v4"
	language "golang.org/x/text/language"
	message "golang.org/x/text/message"
)

type MMPrinter struct{}

func PrintRetail(c echo.Context) error {
	p := MMPrinter{}
	return p.Print(c)
}

func PrintRetailSmall(c echo.Context) error {
	p := MMPrinter{}
	return p.PrintSmall(c)
}

func (m *MMPrinter) center(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

func (m *MMPrinter) Print(c echo.Context) error {

	data := CustomerOrder{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	paperWidth := 48
	pw := 64
	border := strings.Repeat("-", paperWidth)

	s := strings.Builder{}

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString(m.center("TOKO MM", paperWidth))
	s.WriteString(fmt.Sprintf("\n%c%c%c", 27, 77, 1))
	s.WriteString(m.center(fmt.Sprintf("ID: %d - TGL: %s - KASIR: %s", data.ID, data.CreatedAt, data.UpdatedBy), pw))
	s.WriteString(fmt.Sprintf("%c%c%c\n", 27, 77, 0))
	// s.WriteString(m.center(fmt.Sprintf("ORDER ID: %d", data.ID), paperWidth))
	// s.WriteString("\n")
	// s.WriteString(m.center(fmt.Sprintf("TANGGAL: %s", data.CreatedAt), paperWidth))
	// s.WriteString("\n")
	// s.WriteString(m.center(fmt.Sprintf("KASIR: %s", data.UpdatedBy), paperWidth))
	// s.WriteString("\n")
	s.WriteString(m.center(border, paperWidth))
	s.WriteString("\n")

	p := message.NewPrinter(language.Indonesian)
	n := 0
	q := 0.0

	for _, d := range data.Details {
		pot := p.Sprintf("%0.f", d.Pot)
		sqty := fmt.Sprintf("%v %s @ %s", d.Qty, d.Unit, pot)
		nqty := len(sqty)

		name := d.Name
		if len(d.VariantName) > 1 {
			name += ", " + d.VariantName
		}

		maxlen := paperWidth - 10 - nqty
		if len(name) > maxlen {
			name = name[0:maxlen]
		}

		desc := fmt.Sprintf("%s, %s", name, sqty)

		subtotal := p.Sprintf("%0.f", d.Subtotal)
		s.WriteString(fmt.Sprintf("%-38s%10s", desc, subtotal))
		s.WriteString("\n")
		n = n + 1
		q = q + d.Qty
	}

	s.WriteString(m.center(border, paperWidth))
	s.WriteString("\n")

	label := "KEMBALI"
	kembali := data.Payment - data.Total

	item := fmt.Sprintf("Total Item %d / %v", n, q)

	s.WriteString(fmt.Sprintf("%-20s%10s%18s", item, "TOTAL", p.Sprintf("%0.f", data.Total)))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf("%30s%18s", "BAYAR", p.Sprintf("%0.f", data.Payment)))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf("%30s%18s", label, p.Sprintf("%0.f", kembali)))
	s.WriteString("\n")

	log.Printf("%-25s#%v", "Print nota retail:", data.ID)

	m.print(s.String())

	return c.JSON(http.StatusOK, HelloWorld{
		Message: fmt.Sprintf("Print success Retail No. %d", data.ID),
	})
}

func (m *MMPrinter) PrintSmall(c echo.Context) error {

	data := CustomerOrder{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	paperWidth := 64
	pw := 48
	border := strings.Repeat("-", paperWidth)

	s := strings.Builder{}

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString(fmt.Sprintf("%c%c%c", 27, 77, 0))
	s.WriteString(m.center("TOKO MM", pw))
	s.WriteString(fmt.Sprintf("%c%c%c", 27, 77, 1))
	s.WriteString("\n")
	s.WriteString(m.center(fmt.Sprintf("ID: %d - TGL: %s - KASIR: %s", data.ID, data.CreatedAt, data.UpdatedBy), paperWidth))
	s.WriteString("\n")
	// s.WriteString(m.center(fmt.Sprintf("ORDER ID: %d", data.ID), paperWidth))
	// s.WriteString("\n")
	// s.WriteString(m.center(fmt.Sprintf("TANGGAL: %s", data.CreatedAt), paperWidth))
	// s.WriteString("\n")
	// s.WriteString(m.center(fmt.Sprintf("KASIR: %s", data.UpdatedBy), paperWidth))
	// s.WriteString("\n")
	s.WriteString(m.center(border, paperWidth))
	s.WriteString("\n")

	p := message.NewPrinter(language.Indonesian)
	n := 0
	q := 0.0

	for _, d := range data.Details {

		pot := p.Sprintf("%0.f", d.Pot)
		sqty := fmt.Sprintf("%v %s @ %s", d.Qty, d.Unit, pot)
		nqty := len(sqty)

		name := d.Name
		if len(d.VariantName) > 1 {
			name += ", " + d.VariantName
		}

		maxlen := paperWidth - 12 - nqty
		if len(name) > maxlen {
			name = name[0:maxlen]
		}

		desc := fmt.Sprintf("%s, %s", name, sqty)

		subtotal := p.Sprintf("%0.f", d.Subtotal)
		s.WriteString(fmt.Sprintf("%-52s%12s", desc, subtotal))
		s.WriteString("\n")

		n = n + 1
		q = q + d.Qty

	}

	s.WriteString(m.center(border, paperWidth))
	s.WriteString("\n")

	label := "KEMBALI"
	kembali := data.Payment - data.Total

	item := fmt.Sprintf("Total Item %d / %v", n, q)
	s.WriteString(fmt.Sprintf("%-20s%26s%18s", item, "TOTAL", p.Sprintf("%0.f", data.Total)))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf("%46s%18s", "BAYAR", p.Sprintf("%0.f", data.Payment)))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf("%46s%18s", label, p.Sprintf("%0.f", kembali)))
	s.WriteString("\n")

	log.Printf("%-25s#%v", "Print nota retail:", data.ID)

	m.print(s.String())

	return c.JSON(http.StatusOK, HelloWorld{
		Message: fmt.Sprintf("Print success Retail No. %d", data.ID),
	})
}

func (m *MMPrinter) print(data string) {

	//print_logo()

	//	return

	// Ambil nama default printer di windows
	printerName, _ := printer.Default()

	// Membuka printer
	p, err := printer.Open(printerName)

	if err != nil {
		fmt.Println(err)
	}
	// Menutup printer setelah selesai digunakan
	defer p.Close()

	// Memberikan nama/judul document di queue/antrian yang akan di cetak
	err = p.StartRawDocument("Retail")
	if err != nil {
		fmt.Println(err)
	}
	// Menutup document/file setelah selesai digunakan
	defer p.EndDocument()

	// Memulai halaman untuk dicetak
	err = p.StartPage()
	if err != nil {
		fmt.Println(err)
	}
	esc := Esc{}
	esc.Init(p)
	esc.Character(p, 0)
	esc.Print(p, data)
	esc.Feed(p, 4)
	esc.Cut(p)
}
