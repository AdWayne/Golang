package main

func main() {
	pdf := &PDFReport{}
	GenerateReport(pdf)

	excel := &ExcelReport{}
	GenerateReport(excel)

	html := &HTMLReport{}
	GenerateReport(html)

	csv := &CSVReport{}
	GenerateReport(csv)
}