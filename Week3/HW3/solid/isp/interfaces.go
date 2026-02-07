package isp

type Printer interface {
	Print(content string)
}

type Scanner interface {
	Scan(content string)
}

type Fax interface {
	Fax(content string)
}
