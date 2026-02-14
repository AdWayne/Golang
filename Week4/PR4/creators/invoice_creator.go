package creators

import "document-factory/documents"

type InvoiceCreator struct{}

func (c InvoiceCreator) CreateDocument() documents.Document {
	return documents.Invoice{}
}
