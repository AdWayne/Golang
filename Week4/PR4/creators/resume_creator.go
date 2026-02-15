package creators

import "document-factory/documents"

type ResumeCreator struct{}

func (c ResumeCreator) CreateDocument() documents.Document {
	return documents.Resume{}
}
