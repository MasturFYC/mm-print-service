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

func PrintCashier(c echo.Context) error {

	data := CashierReport{}
	if err := c.Bind(&data); err != nil {
		// log.Println(err.Error())
		return err
	}

	s := strings.Builder{}

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString("\u001B\u0057\u0031\u001B\u0045TOKO MM\u001B\u0046\u001B\u0057\u0030\n")
	s.WriteString("LAPORAN PENDAPATAN HARIAN PER USER\n\n")
	s.WriteString("TANGGAL ")
	s.WriteString(data.Created_at)
	s.WriteString("\n\n")
	s.WriteString("------------------------------+-----------------------\n")
	s.WriteString("USER                          |               SUBTOTAL\n")
	s.WriteString("------------------------------+-----------------------\n")

	p := message.NewPrinter(language.Indonesian)
	total := 0.0

	for _, r := range data.Data {
		total += r.Subtotal
		s.WriteString(fmt.Sprintf("%-29s | %22s", r.User, p.Sprintf("%0.f", r.Subtotal)))
		s.WriteString("\n")
	}
	s.WriteString("------------------------------+-----------------------\n")
	s.WriteString(fmt.Sprintf("%-33s%c%c%c%c%14s%c%c%c%c\n", "TOTAL", 27, 80, 27, 69, p.Sprintf("%0.f", total), 27, 103, 27, 70))

	log.Printf("%-25s#%v", "Print report cashir:", data.Created_at)

	print_report_cashier(s.String())
	// log.Printf("%v", s.String())

	return c.JSON(http.StatusOK, HelloWorld{
		Message: fmt.Sprintf("Print report cashier date %s", data.Created_at),
	})
}

func print_report_cashier(data string) {

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
	err = p.StartRawDocument("Laporan harian kasir")
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
	esc := NewEscpos(p)
	esc.Init()
	esc.PageLength(33)
	esc.Pitch(103)
	esc.Typeface(1)
	esc.Print(data)
}
