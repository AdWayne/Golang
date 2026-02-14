package documents

import "fmt"

type Report struct{}

func (r Report) Open() {
	fmt.Println("Открываю документ: Отчет (Report)")
}
