package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alexbrainman/printer"
	"github.com/labstack/echo/v4"
)

//const YYYYMMDD = "2006-01-02"

func PrintDO(c echo.Context) error {

	data := PrintDelivery{}
	if err := c.Bind(&data); err != nil {
		log.Println(err.Error())
		return err
	}

	s := strings.Builder{}

	delivery := data.Delivery
	product := data.Details

	formatStr := "%-35s | %7v | %-8s\n"

	//	tgl, _ := time.Parse(YYYYMMDD, delivery.Created_at)

	//ESC := "\u001B"
	//BOLD := ESC + "\u0045"
	//UNBOLD := ESC + "\u0046"

	s.WriteString(fmt.Sprintf("      \u001B\u0057\u0031\u001B\u0045TOKO MM\u001B\u0046\u001B\u0057\u0030         DO-ID:    %d\n", delivery.Delivery_id))
	s.WriteString(fmt.Sprintf("Jl. Raya Sukra - Indramayu   TANGGAL: %s\n", delivery.Created_at))
	s.WriteString(fmt.Sprintf("  HP/WA: 082 318 321 934     SUPIR:   %s\n", delivery.Driver_name))
	s.WriteString(fmt.Sprintf("                             NOPOL:   %s\n", data.Delivery.Nopol))
	s.WriteString("------------------------------------+---------+-------\n")
	s.WriteString(fmt.Sprintf(formatStr, "NAMA BARANG", "QTY", "UNIT"))
	s.WriteString("------------------------------------+---------+-------\n")

	for _, d := range product {
		name := d.Name
		if d.Variant_name != "" {
			if len(d.Variant_name) > 1 {
				name += ", " + d.Variant_name
			}

			if len(name) > 35 {
				name = name[0:35]
			}
		}

		s.WriteString(fmt.Sprintf(formatStr, name, d.Qty, d.Unit))
	}

	item := ""

	if len(product) > 1 {
		item = "s"
	}

	s.WriteString("------------------------------------+---------+-------\n")
	s.WriteString(fmt.Sprintf("Total item: %v item%s\n", len(product), item))

	log.Printf("%-25s#%v", "Print delivery order:", delivery.Delivery_id)

	print_do(s.String())

	//log.Println(s.String())

	return c.JSON(http.StatusOK, HelloWorld{
		Message: fmt.Sprintf("Print success delivery order No. %d", delivery.Delivery_id),
	})

}

func print_do(data string) {

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
	err = p.StartRawDocument("Faktur DO")
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
