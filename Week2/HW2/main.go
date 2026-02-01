package main

import (
	"errors"
	"fmt"
	"os"
)

// DRY — ЛОГГЕР

type LogLevel string

const (
	ERROR   LogLevel = "ERROR"
	WARNING LogLevel = "WARNING"
	INFO    LogLevel = "INFO"
)

func Log(level LogLevel, message string) {
	fmt.Printf("%s: %s\n", level, message)
}

// DRY — ОБЩАЯ КОНФИГУРАЦИЯ

type Config struct {
	ConnectionString string
}

type DatabaseService struct {
	Config *Config
}

func (d *DatabaseService) Connect() {
	Log(INFO, "Connecting to DB: "+d.Config.ConnectionString)
}

type LoggingService struct {
	Config *Config
}

func (l *LoggingService) SaveLog(message string) {
	Log(INFO, "Saving log to DB: "+l.Config.ConnectionString+" | "+message)
}

// KISS — ПРОСТАЯ ЛОГИКА

func ProcessNumbers(numbers []int) {
	if len(numbers) == 0 {
		return
	}

	for _, n := range numbers {
		if n > 0 {
			fmt.Println("Positive:", n)
		}
	}
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// YAGNI — USER ТОЛЬКО ДАННЫЕ

type User struct {
	Name  string
	Email string
}

type UserRepository struct{}

func (r *UserRepository) Save(user User) {
	Log(INFO, "User saved: "+user.Name)
}

type EmailService struct{}

func (e *EmailService) SendEmail(user User) {
	Log(INFO, "Email sent to: "+user.Email)
}

// YAGNI — ПРОСТОЙ FILE READER

type FileReader struct{}

func (f *FileReader) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// YAGNI — ОДИН ГЕНЕРАТОР ОТЧЁТОВ

type ReportGenerator struct{}

func (r *ReportGenerator) Generate(reportType string) {
	switch reportType {
	case "pdf":
		Log(INFO, "Generating PDF report")
	case "excel":
		Log(INFO, "Generating Excel report")
	case "html":
		Log(INFO, "Generating HTML report")
	default:
		Log(WARNING, "Unknown report type")
	}
}

// MAIN

func main() {
	// DRY config
	cfg := &Config{
		ConnectionString: "Server=myServer;Database=myDb;User=myUser;Password=myPass;",
	}

	db := DatabaseService{Config: cfg}
	logger := LoggingService{Config: cfg}

	db.Connect()
	logger.SaveLog("Application started")

	// KISS numbers
	ProcessNumbers([]int{-2, 5, 0, 9})

	result, err := Divide(10, 2)
	if err != nil {
		Log(ERROR, err.Error())
	} else {
		fmt.Println("Division result:", result)
	}

	// YAGNI user logic
	user := User{Name: "Azamat", Email: "azamat@mail.com"}

	repo := UserRepository{}
	emailService := EmailService{}

	repo.Save(user)
	emailService.SendEmail(user)

	// YAGNI report
	reportGen := ReportGenerator{}
	reportGen.Generate("pdf")
}
