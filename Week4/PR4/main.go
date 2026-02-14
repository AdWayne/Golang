package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const folderName = "documents"

// ================== DOCUMENT INTERFACE ==================
type Document interface {
	Open()
	Save()
	GetFileName() string
}

// ================== BASE FUNCTION ==================
func ensureFolder() {
	os.MkdirAll(folderName, os.ModePerm)
}

func openFile(path string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("notepad", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}

	cmd.Start()
}

// ================== REPORT ==================
type Report struct{}

func (r Report) GetFileName() string {
	return "report.txt"
}

func (r Report) Save() {
	content := "Это документ: Report\nДанные отчета..."
	path := filepath.Join(folderName, r.GetFileName())
	os.WriteFile(path, []byte(content), 0644)
}

func (r Report) Open() {
	path := filepath.Join(folderName, r.GetFileName())
	openFile(path)
}

// ================== RESUME ==================
type Resume struct{}

func (r Resume) GetFileName() string {
	return "resume.txt"
}

func (r Resume) Save() {
	content := "Это документ: Resume\nИмя: Nygmet Azamat"
	path := filepath.Join(folderName, r.GetFileName())
	os.WriteFile(path, []byte(content), 0644)
}

func (r Resume) Open() {
	path := filepath.Join(folderName, r.GetFileName())
	openFile(path)
}

// ================== LETTER ==================
type Letter struct{}

func (l Letter) GetFileName() string {
	return "letter.txt"
}

func (l Letter) Save() {
	content := "Это документ: Letter\nЗдравствуйте!"
	path := filepath.Join(folderName, l.GetFileName())
	os.WriteFile(path, []byte(content), 0644)
}

func (l Letter) Open() {
	path := filepath.Join(folderName, l.GetFileName())
	openFile(path)
}

// ================== INVOICE ==================
type Invoice struct{}

func (i Invoice) GetFileName() string {
	return "invoice.txt"
}

func (i Invoice) Save() {
	content := "Это документ: Invoice\nСумма: 1000$"
	path := filepath.Join(folderName, i.GetFileName())
	os.WriteFile(path, []byte(content), 0644)
}

func (i Invoice) Open() {
	path := filepath.Join(folderName, i.GetFileName())
	openFile(path)
}

// ================== FACTORY ==================
type DocumentCreator interface {
	CreateDocument() Document
}

type ReportCreator struct{}
func (c ReportCreator) CreateDocument() Document { return Report{} }

type ResumeCreator struct{}
func (c ResumeCreator) CreateDocument() Document { return Resume{} }

type LetterCreator struct{}
func (c LetterCreator) CreateDocument() Document { return Letter{} }

type InvoiceCreator struct{}
func (c InvoiceCreator) CreateDocument() Document { return Invoice{} }

// ================== GET CREATOR ==================
func getCreator(docType string) DocumentCreator {
	switch docType {
	case "report":
		return ReportCreator{}
	case "resume":
		return ResumeCreator{}
	case "letter":
		return LetterCreator{}
	case "invoice":
		return InvoiceCreator{}
	default:
		return nil
	}
}

// ================== MAIN ==================
func main() {

	ensureFolder()

	fmt.Println("Введите тип документа (report, resume, letter, invoice):")
	fmt.Println("Если документ уже создан — он откроется.")
	fmt.Println("Если нет — он создастся и откроется.")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	creator := getCreator(input)
	if creator == nil {
		fmt.Println("Неизвестный тип документа")
		return
	}

	doc := creator.CreateDocument()
	doc.Save()
	doc.Open()

	fmt.Println("Документ открыт из папки:", folderName)
}
