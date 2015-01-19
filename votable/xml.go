package votable

import (
	"encoding/xml"
)

type xmlVO struct {
	XMLName   xml.Name      `xml:"VOTABLE"`
	Version   string        `xml:"version,attr"`
	Resources []xmlResource `xml:"RESOURCE"`
}

type xmlResource struct {
	Name   string     `xml:"name,attr"`
	Tables []xmlTable `xml:"TABLE"`
}

type xmlTable struct {
	Name        string       `xml:"name,attr"`
	Description string       `xml:"DESCRIPTION"`
	Groups      []xmlGroup   `xml:"GROUP"`
	Param       xmlParam     `xml:"PARAM"`
	Fields      []xmlField   `xml:"FIELD"`
	TableData   xmlTableData `xml:"DATA>TABLEDATA"`
	BinTable    xmlBinData   `xml:"DATA>BINARY"`
	BinTable2   xmlBinData2  `xml:"DATA>BINARY2"`
}

type xmlGroup struct {
	ID        string        `xml:"ID,attr"`
	UType     string        `xml:"utype,attr"`
	Params    []xmlParam    `xml:"PARAM"`
	FieldRefs []xmlFieldRef `xml:"FIELDref"`
}

type xmlParam struct {
	Name      string `xml:"name,attr"`
	DataType  string `xml:"datatype,attr"`
	ArraySize string `xml:"arraysize,attr"`
	UCD       string `xml:"ucd,attr"`
	UType     string `xml:"utype,attr"`
	Value     string `xml:"value,attr"`
}

type xmlFieldRef struct {
	Ref string `xml:"ref,attr"`
}

type xmlField struct {
	Name        string  `xml:"name,attr"`
	ID          string  `xml:"ID,attr"`
	UCD         string  `xml:"ucd,attr"`
	Ref         string  `xml:"ref,attr"`
	UType       string  `xml:"utype,attr"`
	DataType    string  `xml:"datatype,attr"`
	ArraySize   string  `xml:"arraysize,attr"`
	Width       int     `xml:"width,attr"`
	Precision   int     `xml:"precision,attr"`
	Unit        string  `xml:"unit,attr"`
	Description string  `xml:"DESCRIPTION"`
	Link        xmlLink `xml:"LINK,omitempty"`
}

type xmlLink struct {
	Type string `xml:"type,attr"`
	HRef string `xml:"href,attr"`
}

type xmlRow struct {
	Data [][]byte `xml:"TD"`
}

type xmlTableData struct {
	Rows []xmlRow `xml:"TR"`
}

type xmlBinData struct{}
type xmlBinData2 struct{}
