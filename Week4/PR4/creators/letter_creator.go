package creators

import "document-factory/documents"

type LetterCreator struct{}

func (c LetterCreator) CreateDocument() documents.Document {
	return documents.Letter{}
}
