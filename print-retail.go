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

func (m *MMPrinter) center(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

func (m *MMPrinter) Print(c echo.Context) error {

	data := CustomerOrder{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	paperWidth := 110
	border := strings.Repeat("-", paperWidth)

	s := strings.Builder{}

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString(m.center("\u001B\u0057\u0031\u001B\u0045TOKO MM\u001B\u0046\u001B\u0057\u0030", paperWidth))
	s.WriteString("\n")
	s.WriteString(m.center(fmt.Sprintf("ORDER ID: %d", data.ID), paperWidth))
	s.WriteString("\n")
	s.WriteString(m.center(fmt.Sprintf("TANGGAL: %s", data.CreatedAt), paperWidth))
	s.WriteString("\n")
	s.WriteString(m.center(fmt.Sprintf("KASIR: %s", data.UpdatedBy), paperWidth))
	s.WriteString("\n")
	s.WriteString(m.center(border, paperWidth))
	s.WriteString("\n")

	p := message.NewPrinter(language.Indonesian)

	for _, d := range data.Details {
		name := d.Name
		if len(d.VariantName) > 1 {
			name += ", " + d.VariantName
		}

		if len(name) > 17 {
			name = name[0:17]
		}

		s.WriteString(fmt.Sprintf("%-17s | %4v %-4s | %9s | %10s", name, d.Qty, d.Unit,
			p.Sprintf("%0.f", d.Pot), p.Sprintf("%0.f", d.Subtotal)))
		s.WriteString("\n")
	}

	s.WriteString(m.center(border, paperWidth))
	s.WriteString("\n")

	//s.WriteString(fmt.Sprintf("%20s%10s%c%c%c%12s%c%c%c\n", "", "TOTAL", 27, 87, 1, p.Sprintf("%0.f", data.Total), 27, 87, 0))
	s.WriteString(fmt.Sprintf("%26s%10s%c%c%c%c%12s%c%c%c%c", "", "TOTAL", 27, 80, 27, 69, p.Sprintf("%0.f", data.Total), 27, 103, 27, 70))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf("%26s%10s%18s", "", "BAYAR", p.Sprintf("%0.f", data.Payment)))
	s.WriteString("\n")

	label := "KEMBALI"
	kembali := data.Payment - data.Total

	s.WriteString(fmt.Sprintf("%26s%10s%18s", "", label, p.Sprintf("%0.f", kembali)))
	s.WriteString("\n")

	log.Printf("%-25s#%v", "Print nota retail:", data.ID)

	m.print(s.String())
	// log.Printf("%v", s.String())

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
	err = p.StartRawDocument("Faktur")
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
	esc.PageLength(p, 33)
	esc.Pitch(p, 103)
	esc.Typeface(p, 1)
	esc.Print(p, data)
}
