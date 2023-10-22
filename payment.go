package main

import (
	"fmt"
	"log"
	"net/http"

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
	esc := Esc{}
	esc.Init(p)
	esc.PageLength(p, 33)
	esc.Pitch(p, 80)
	esc.Typeface(p, 1)
	esc.Print(p, "\u001B\u0057\u0031\u001B\u0045KWITANSI\u001B\u0046\u001B\u0057\u0030")
	esc.Print(p, "\n")
	esc.Print(p, "\n")
	esc.Print(p, "\n")
	esc.Print(p, fmt.Sprintf("%-17s : %v\n", "NO. Kwitansi", data.PaymentId))
	esc.Print(p, fmt.Sprintf("%-17s : %s\n\n", "Telah terima dari", data.CustomerName))
	esc.Print(p, fmt.Sprintf("%-17s :\n", "Uang sejumlah"))
	esc.Pitch(p, 77)
	esc.Italic(p)
	esc.Print(p, fmt.Sprintf("\"%s%s\"\n\n", data.Terbilang, "Rupiah"))
	esc.Unitalic(p)
	esc.Pitch(p, 80)
	esc.Print(p, fmt.Sprintf("%-17s :\n%s #%v\n\n", "Untuk pembayaran", "Piutang No. Order", data.OrderId))
	esc.Print(p, fmt.Sprintf("%-17s : Rp.", "Terbilang"))
	esc.Bold(p)
	esc.Print(p, f.Sprintf("%0.f", data.Amount))
	esc.Unbold(p)
	esc.Print(p, "\n\n")
	esc.Print(p, "\n")
	esc.Pitch(p, 77)
	esc.Print(p, fmt.Sprintf("%-15sIndramayu, %s\n", "", data.CreatedAt))
	esc.Print(p, fmt.Sprintf("%-15s%s\n", "", "a.n. Admin Toko MM,"))
	esc.Print(p, "\n\n")
	esc.Print(p, fmt.Sprintf("%-15s%s\n", "", data.Admin))
	esc.Print(p, "\n")
	//	esc.Print(p, "───")
}
