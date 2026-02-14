package creators

import "document-factory/documents"

type ReportCreator struct{}

func (c ReportCreator) CreateDocument() documents.Document {
	return documents.Report{}
}
