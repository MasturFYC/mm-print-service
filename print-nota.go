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

func PrintNota(c echo.Context) error {

	data := CustomerOrder{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	s := strings.Builder{}

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString(fmt.Sprintf("      \u001B\u0057\u0031\u001B\u0045TOKO MM\u001B\u0046\u001B\u0057\u0030        ORDER ID: %d\n", data.ID))
	s.WriteString(fmt.Sprintf("Jl. Raya Sukra - Indramayu  SALES:    %s\n", data.SalesName))
	s.WriteString(fmt.Sprintf("  HP/WA: 082 318 321 934    PLGN:     %s\n", data.CustomerName))
	s.WriteString(fmt.Sprintf("                            ALAMAT:   %s\n", data.Address))
	s.WriteString(fmt.Sprintf("                            TANGGAL:  %s\n", data.CreatedAt))
	s.WriteString(fmt.Sprintf("                            USER:     %s\n", data.UpdatedBy))
	s.WriteString("------------------+-----------+-----------+-----------\n")
	s.WriteString("NAMA BARANG       |  QTY UNIT |     HARGA |   SUBTOTAL\n")
	s.WriteString("------------------+-----------+-----------+-----------\n")

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

	s.WriteString("------------------+-----------+-----------+-----------\n")
	//s.WriteString(fmt.Sprintf("%20s%10s%c%c%c%12s%c%c%c\n", "", "TOTAL", 27, 87, 1, p.Sprintf("%0.f", data.Total), 27, 87, 0))
	s.WriteString(fmt.Sprintf("%26s%10s%c%c%c%c%12s%c%c%c%c\n", "", "TOTAL", 27, 80, 27, 69, p.Sprintf("%0.f", data.Total), 27, 103, 27, 70))
	s.WriteString(fmt.Sprintf("%26s%10s%18s\n", "", "BAYAR", p.Sprintf("%0.f", data.Payment)))

	label := "KEMBALI"
	kembali := data.Payment - data.Total

	if data.Payment < data.Total {
		label = "PIUTANG"
		kembali = data.Total - data.Payment
	}

	s.WriteString(fmt.Sprintf("%26s%10s%18s\n", "", label, p.Sprintf("%0.f", kembali)))

	log.Printf("%-25s#%v", "Print nota order:", data.ID)

	print_nota(s.String())
	// log.Printf("%v", s.String())

	return c.JSON(http.StatusOK, HelloWorld{
		Message: fmt.Sprintf("Print success Nota No. %d", data.ID),
	})
}

func print_nota(data string) {

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
	esc := NewEscpos(p)
	esc.Init()
	esc.PageLength(33)
	esc.Pitch(103)
	esc.Typeface(1)
	esc.Print(data)
}

// func print_logo() {

// 	printerName, _ := goprint.GetDefaultPrinterName()

// 	//open the printer
// 	printerHandle, err := goprint.GoOpenPrinter(printerName)
// 	if err != nil {
// 		log.Fatalln("Failed to open printer")
// 	}
// 	defer goprint.GoClosePrinter(printerHandle)

// 	filePath := "C:/Users/mastu/Documents/godoc/tokomm.pdf"

// 	//Send to printer:
// 	err = goprint.GoPrint(printerHandle, filePath)
// 	if err != nil {
// 		log.Fatalln("during the func sendToPrinter, there was an error")
// 	}
// }
