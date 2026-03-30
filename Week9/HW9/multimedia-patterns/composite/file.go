package composite

import "fmt"

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name, size}
}

func (f *File) Display(indent string) {
	fmt.Println(indent + "- " + f.name)
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) GetName() string {
	return f.name
}