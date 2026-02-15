package documents

import "fmt"

type Invoice struct{}

func (i Invoice) Open() {
	fmt.Println("Открываю документ: Счет (Invoice)")
}
