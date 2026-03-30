package composite

type FileSystemComponent interface {
	Display(indent string)
	GetSize() int
	GetName() string
}