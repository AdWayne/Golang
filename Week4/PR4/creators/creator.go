package creators

import "document-factory/documents"

// DocumentCreator — абстрактный создатель
type DocumentCreator interface {
	CreateDocument() documents.Document
}
