package documents

import "fmt"

type Letter struct{}

func (l Letter) Open() {
	fmt.Println("Открываю документ: Письмо (Letter)")
}
