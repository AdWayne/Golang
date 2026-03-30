package composite

import "fmt"

type Directory struct {
	name       string
	components []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

// добавление
func (d *Directory) Add(component FileSystemComponent) {
	for _, c := range d.components {
		if c.GetName() == component.GetName() {
			fmt.Println("Уже существует:", component.GetName())
			return
		}
	}

	d.components = append(d.components, component)
}

// удаление
func (d *Directory) Remove(name string) {
	for i, c := range d.components {
		if c.GetName() == name {
			d.components = append(d.components[:i], d.components[i+1:]...)
			return
		}
	}

	fmt.Println("Не найден:", name)
}

func (d *Directory) Display(indent string) {
	fmt.Println(indent + "[DIR] " + d.name)

	for _, c := range d.components {
		c.Display(indent + "  ")
	}
}

func (d *Directory) GetSize() int {
	total := 0
	for _, c := range d.components {
		total += c.GetSize()
	}
	return total
}

func (d *Directory) GetName() string {
	return d.name
}