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
	"github.com/jadefox10200/goprint"
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

	sep := "-"
	sepCount := 20 + 5 + 6 + 9 + 12 + 1
	sep2 := "%39s"
	sep3 := "%15s"
	s1 := "%-20v"
	s2 := "%5v"
	s3 := " %-6v"
	s4 := "%9s"
	s5 := "%12s"

	s := strings.Builder{}

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString(fmt.Sprintf("      \u001B\u0057\u0031\u001B\u0045TOKO MM\u001B\u0046\u001B\u0057\u0030         ORDER ID:  %d\n", data.ID))
	s.WriteString(fmt.Sprintf("Jl. Raya Sukra - Indramayu   SALES:     %s\n", data.SalesName))
	s.WriteString(fmt.Sprintf("  HP/WA: 082 318 321 934     PELANGGAN: %s\n", data.CustomerName))
	s.WriteString(fmt.Sprintf("                             ALAMAT:    %s\n", data.Address))
	s.WriteString(strings.Repeat(sep, sepCount))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf(s1, "NAMA BARANG"))
	s.WriteString(fmt.Sprintf(s2, "QTY"))
	s.WriteString(fmt.Sprintf(s3, "UNIT"))
	s.WriteString(fmt.Sprintf(s4, "HARGA"))
	s.WriteString(fmt.Sprintf(s5, "SUBTOTAL"))
	s.WriteString("\n")
	s.WriteString(strings.Repeat(sep, sepCount))
	s.WriteString("\n")
	p := message.NewPrinter(language.Indonesian)

	for _, d := range data.Details {
		name := d.Name
		if len(d.VariantName) > 1 {
			name += ", " + d.VariantName
		}

		if len(name) > 20 {
			name = name[0:20]
		}

		s.WriteString(fmt.Sprintf(s1, name))
		s.WriteString(fmt.Sprintf(s2, d.Qty))
		s.WriteString(fmt.Sprintf(s3, d.Unit))
		s.WriteString(fmt.Sprintf(s4, p.Sprintf("%0.f", d.Pot)))
		s.WriteString(fmt.Sprintf(s5, p.Sprintf("%0.f", d.Subtotal)))
		s.WriteString("\n")
	}
	s.WriteString(strings.Repeat(sep, sepCount))
	s.WriteString("\n")
	s.WriteString(fmt.Sprintf(sep2, "TOTAL"))
	s.WriteString(fmt.Sprintf("%19s", "\u001B\u0050\u001B\u0045"+p.Sprintf("%0.f\u001B\u0046\u001B\u0067\n", data.Total)))
	s.WriteString(fmt.Sprintf(sep2, "BAYAR"))
	s.WriteString(fmt.Sprintf(sep3, p.Sprintf("%0.f\n", data.Payment)))
	s.WriteString(fmt.Sprintf(sep2, "SISA BAYAR"))
	s.WriteString(fmt.Sprintf(sep3, p.Sprintf("%0.f\n", data.RemainPayment)))

	log.Printf("%-25s#%v", "Print nota order:", data.ID)

	print_nota(s.String())

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

	// **************** LIHAT DI CHARACTER MAP ***************** //

	// Mulai mencetak string ke printer default yang ada di windows
	ESC := "\u001B"
	//CONDENSED := ESC + "\u0021\u0004" // "\u0065"
	TYPEFACE := ESC + "\u006B"
	SANS := TYPEFACE + "\u0031"
	//PITCH10 := ESC + "\u0050"
	//PITCH12 := ESC + "\u004D"
	PITCH15 := ESC + "\u0067"
	//DRAFT := ESC + "\u0078\u0030"
	//OCRB := TYPEFACE + "\u0005"
	//OCRA := TYPEFACE + "\u0006"
	//ORATOR := TYPEFACE + "\u0007"
	//SANSH := TYPEFACE + "\u0011"
	fmt.Fprint(p, ESC+"@")
	//fmt.Fprint(p, DRAFT)
	//fmt.Fprint(p, CONDENSED)
	fmt.Fprint(p, PITCH15)
	fmt.Fprint(p, SANS)
	fmt.Fprint(p, ESC+"\u0012")
	fmt.Fprint(p, data)
}

func print_logo() {

	printerName, _ := goprint.GetDefaultPrinterName()

	//open the printer
	printerHandle, err := goprint.GoOpenPrinter(printerName)
	if err != nil {
		log.Fatalln("Failed to open printer")
	}
	defer goprint.GoClosePrinter(printerHandle)

	filePath := "C:/Users/mastu/Documents/godoc/tokomm.pdf"

	//Send to printer:
	err = goprint.GoPrint(printerHandle, filePath)
	if err != nil {
		log.Fatalln("during the func sendToPrinter, there was an error")
	}
}
