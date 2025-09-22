package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alexbrainman/printer"
	"github.com/labstack/echo/v4"
	language "golang.org/x/text/language"
	message "golang.org/x/text/message"
)

//const YYYYMMDD = "2006-01-02"

func PrintPayment(c echo.Context) error {

	data := PrintDataPayment{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	print_payment(data)

	log.Printf("%-25s#%v", "Print order payment:", data.PaymentId)

	return c.JSON(http.StatusOK, HelloWorld{
		Message: fmt.Sprintf("Print piutang suceess %d", data.PaymentId),
	})

}

func print_payment(data PrintDataPayment) {

	//print_logo()

	//	return
	f := message.NewPrinter(language.Indonesian)

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
	err = p.StartRawDocument("Kwitansi Piutang")
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

	// **************** LIHAT DI CHARACTER MAP ***************** //

	// Mulai mencetak string ke printer default yang ada di windows
	esc := NewEscpos(p)
	esc.Init()
	esc.PageLength(33)
	esc.Pitch(80)
	esc.Typeface(0)
	esc.LqMode(1)
	esc.Print("\u001B\u0057\u0031\u001B\u0045KWITANSI\u001B\u0046\u001B\u0057\u0030")
	esc.Print("\n")
	esc.Print("\n")
	esc.Print("\n")
	esc.Print(fmt.Sprintf("Kwitansi No.      : %v\n\n", data.PaymentId))
	esc.Print(fmt.Sprintf("Telah terima dari : %s\n\n", data.CustomerName))
	esc.Print("Uang sejumlah     :\n")
	esc.Print("\"")
	esc.Italic()
	esc.Print(break_text(data.Terbilang))
	esc.Unitalic()
	esc.Print("\"\n\n")
	esc.Print(fmt.Sprintf("Untuk pembayaran  :\n%s #%v\n\n", "Piutang No. Order", data.OrderId))
	esc.Print("\n")
	esc.Print("Terbilang         : Rp.")
	esc.Bold()
	esc.Print(f.Sprintf("%0.f", data.Amount))
	esc.Unbold()
	esc.Print("\n\n")
	esc.Print("\n")
	esc.Pitch(77)
	esc.Print(fmt.Sprintf("%-15sIndramayu, %s\n", "", data.CreatedAt))
	esc.Print(fmt.Sprintf("%-15s%s\n", "", "a.n. Admin Toko MM,"))
	esc.Print("\n\n")
	esc.Print(fmt.Sprintf("%-15s%s\n", "", data.Admin))
	esc.Print("\n")
	//	esc.Print(p, "───")
	// log.Println(break_text(data.Terbilang))
}

func break_text(s string) string {
	ss := strings.Split(s, " ")
	sr := make([]string, len(ss)+2)

	l := 0

	for _, t := range ss {
		l += len(t)
		if l >= 30 {
			sr = append(sr, "\n")
			l = len(t)
		}
		sr = append(sr, t)
	}

	sr = append(sr, "Rupiah")

	return strings.TrimSpace(strings.Join(sr, " "))
}
