package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

////////////////////////////////////////////////////////////
//////////////////////// SINGLETON /////////////////////////
////////////////////////////////////////////////////////////

type ConfigurationManager struct {
	settings map[string]string
}

var singletonInstance *ConfigurationManager
var once sync.Once

func GetConfigurationManager() *ConfigurationManager {
	once.Do(func() {
		singletonInstance = &ConfigurationManager{
			settings: make(map[string]string),
		}
		fmt.Println("ConfigurationManager created")
	})
	return singletonInstance
}

func (c *ConfigurationManager) Set(key, value string) {
	c.settings[key] = value
}

func (c *ConfigurationManager) Get(key string) (string, error) {
	val, ok := c.settings[key]
	if !ok {
		return "", fmt.Errorf("setting '%s' not found", key)
	}
	return val, nil
}

func (c *ConfigurationManager) SaveToFile(filename string) error {
	var builder strings.Builder
	for k, v := range c.settings {
		builder.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}
	return os.WriteFile(filename, []byte(builder.String()), 0644)
}

func (c *ConfigurationManager) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			c.settings[parts[0]] = parts[1]
		}
	}
	return nil
}

func testSingleton() {
	fmt.Println("===== SINGLETON TEST =====")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			config := GetConfigurationManager()
			config.Set(fmt.Sprintf("key%d", id), fmt.Sprintf("value%d", id))
			fmt.Println("Goroutine", id, "instance:", config)
		}(i)
	}

	wg.Wait()
}

////////////////////////////////////////////////////////////
//////////////////////// BUILDER ///////////////////////////
////////////////////////////////////////////////////////////

type Report struct {
	Header  string
	Content string
	Footer  string
}

func (r *Report) Show() {
	fmt.Println(r.Header)
	fmt.Println(r.Content)
	fmt.Println(r.Footer)
	fmt.Println("----------------------")
}

type IReportBuilder interface {
	SetHeader(string)
	SetContent(string)
	SetFooter(string)
	GetReport() *Report
}

//// Text Builder

type TextReportBuilder struct {
	report *Report
}

func NewTextReportBuilder() *TextReportBuilder {
	return &TextReportBuilder{report: &Report{}}
}

func (b *TextReportBuilder) SetHeader(h string) {
	b.report.Header = "TEXT HEADER: " + h
}

func (b *TextReportBuilder) SetContent(c string) {
	b.report.Content = "TEXT CONTENT: " + c
}

func (b *TextReportBuilder) SetFooter(f string) {
	b.report.Footer = "TEXT FOOTER: " + f
}

func (b *TextReportBuilder) GetReport() *Report {
	return b.report
}

//// HTML Builder

type HtmlReportBuilder struct {
	report *Report
}

func NewHtmlReportBuilder() *HtmlReportBuilder {
	return &HtmlReportBuilder{report: &Report{}}
}

func (b *HtmlReportBuilder) SetHeader(h string) {
	b.report.Header = "<h1>" + h + "</h1>"
}

func (b *HtmlReportBuilder) SetContent(c string) {
	b.report.Content = "<p>" + c + "</p>"
}

func (b *HtmlReportBuilder) SetFooter(f string) {
	b.report.Footer = "<footer>" + f + "</footer>"
}

func (b *HtmlReportBuilder) GetReport() *Report {
	return b.report
}

//// Director

type ReportDirector struct{}

func (d *ReportDirector) Construct(builder IReportBuilder) {
	builder.SetHeader("Monthly Report")
	builder.SetContent("Sales analytics")
	builder.SetFooter("Company Confidential")
}

func testBuilder() {
	fmt.Println("\n===== BUILDER TEST =====")

	director := &ReportDirector{}

	textBuilder := NewTextReportBuilder()
	director.Construct(textBuilder)
	textBuilder.GetReport().Show()

	htmlBuilder := NewHtmlReportBuilder()
	director.Construct(htmlBuilder)
	htmlBuilder.GetReport().Show()
}

////////////////////////////////////////////////////////////
//////////////////////// PROTOTYPE /////////////////////////
////////////////////////////////////////////////////////////

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func (p *Product) Clone() *Product {
	return &Product{
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
}

type Discount struct {
	Description string
	Amount      float64
}

func (d *Discount) Clone() *Discount {
	return &Discount{
		Description: d.Description,
		Amount:      d.Amount,
	}
}

type Order struct {
	Products      []*Product
	Discounts     []*Discount
	ShippingCost  float64
	PaymentMethod string
}

func (o *Order) Clone() *Order {
	clone := &Order{
		ShippingCost:  o.ShippingCost,
		PaymentMethod: o.PaymentMethod,
	}

	for _, p := range o.Products {
		clone.Products = append(clone.Products, p.Clone())
	}

	for _, d := range o.Discounts {
		clone.Discounts = append(clone.Discounts, d.Clone())
	}

	return clone
}

func (o *Order) Show() {
	fmt.Println("Order:")
	for _, p := range o.Products {
		fmt.Printf("- %s x%d (%.2f)\n", p.Name, p.Quantity, p.Price)
	}
	for _, d := range o.Discounts {
		fmt.Printf("- Discount: %s (%.2f)\n", d.Description, d.Amount)
	}
	fmt.Println("Shipping:", o.ShippingCost)
	fmt.Println("Payment:", o.PaymentMethod)
	fmt.Println("----------------------")
}

func testPrototype() {
	fmt.Println("\n===== PROTOTYPE TEST =====")

	order1 := &Order{
		Products: []*Product{
			{Name: "Laptop", Price: 1200, Quantity: 1},
			{Name: "Mouse", Price: 30, Quantity: 2},
		},
		Discounts: []*Discount{
			{Description: "Black Friday", Amount: 100},
		},
		ShippingCost:  15,
		PaymentMethod: "Credit Card",
	}

	order2 := order1.Clone()
	order2.Products[0].Quantity = 2
	order2.PaymentMethod = "PayPal"

	fmt.Println("Original:")
	order1.Show()

	fmt.Println("Cloned:")
	order2.Show()
}

////////////////////////////////////////////////////////////
//////////////////////////// MAIN //////////////////////////
////////////////////////////////////////////////////////////

func main() {
	testSingleton()
	testBuilder()
	testPrototype()
}