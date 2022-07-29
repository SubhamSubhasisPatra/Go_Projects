package main

import (
	"bytes"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func CustomError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	CustomError(err)

	pdfg.NoCollate.Set(false)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	// html1, _ := os.ReadFile("./index.ejs") // or the html as a string
	// html2, _ := os.ReadFile("./index.ejs")
	// html3, _ := os.ReadFile("./index.ejs")

	// refer the above code to read from the file system 

	html1 := `<html><head></head><body>HELLO PDF PAGE 1</body></html>`
	html2 := `<html><head></head><body>HELLO PDF PAGE 2</body></html>`
	html3 := `<html><head></head><body>HELLO PDF PAGE 3</body></html>`

	arr := []string{string(html1), string(html2), string(html3)}

	buf := new(bytes.Buffer)

	// itterate over the array and add each page to the pdf
	for _, html := range arr {
		buf.WriteString(html)
		// If you want a page braker for each page add the bellow code , else skip this part 
		buf.WriteString(`<P style="page-break-before: always">`) 
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(buf))

	err = pdfg.Create()

	CustomError(err)

	err = pdfg.WriteFile("simplesample_Multi_page.pdf")

	CustomError(err)
}
