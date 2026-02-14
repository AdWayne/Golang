package documents

import "fmt"

type Resume struct{}

func (r Resume) Open() {
	fmt.Println("Открываю документ: Резюме (Resume)")
}
