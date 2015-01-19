package votable

import (
	"encoding/xml"
)

type xmlVO struct {
	XMLName     xml.Name      `xml:"VOTABLE"`
	Version     string        `xml:"version,attr"`
	Resources   []xmlResource `xml:"RESOURCE"`
	Description string        `xml:"DESCRIPTION,omitempty"`
	Infos       []xmlInfo     `xml:"INFO,omitempty"`
}

type xmlResource struct {
	Name  string `xml:"name,attr,omitempty"`
	ID    string `xml:"ID,attr,omitempty"`
	Type  string `xml:"type,attr,omitempty"`
	UType string `xml:"utype,attr,omitempty"`

	Description string     `xml:"DESCRIPTION,omitempty"`
	Infos       []xmlInfo  `xml:"INFO,omitempty"`
	Tables      []xmlTable `xml:"TABLE"`
}

type xmlInfo struct {
	XMLName xml.Name `xml:"INFO"`
	ID      string   `xml:"ID,attr,omitempty"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
	Unit    string   `xml:"unit,attr,omitempty"`
	XType   string   `xml:"xtype,attr,omitempty"`
	Ref     string   `xml:"ref,attr,omitempty"`
	UCD     string   `xml:"ucd,attr,omitempty"`
	UType   string   `xml:"utype,attr,omitempty"`
}

type xmlTable struct {
	Name        string        `xml:"name,attr"`
	Description string        `xml:"DESCRIPTION"`
	Groups      []xmlGroup    `xml:"GROUP"`
	Param       xmlParam      `xml:"PARAM"`
	Fields      []xmlField    `xml:"FIELD"`
	TableData   *xmlTableData `xml:"DATA>TABLEDATA,omitempty"`
	BinTable    *xmlBinData   `xml:"DATA>BINARY,omitempty"`
	BinTable2   *xmlBinData2  `xml:"DATA>BINARY2,omitempty"`
}

type xmlGroup struct {
	ID    string `xml:"ID,attr,omitempty"`
	Name  string `xml:"name,attr,omitempty"`
	Ref   string `xml:"ref,attr,omitempty"`
	UCD   string `xml:"ucd,attr,omitempty"`
	UType string `xml:"utype,attr,omitempty"`

	Params    []xmlParam    `xml:"PARAM"`
	Groups    []xmlGroup    `xml:"GROUP"`
	ParamRefs []xmlParamRef `xml:"PARAMref"`
	FieldRefs []xmlFieldRef `xml:"FIELDref"`
}

type xmlParamRef struct {
	Ref   string `xml:"ref,attr"`
	UCD   string `xml:"ucd,attr,omitempty"`
	UType string `xml:"utype,attr,omitempty"`
}

type xmlParam struct {
	Name      string `xml:"name,attr"`
	ID        string `xml:"ID,attr,omitempty"`
	Unit      string `xml:"unit,attr,omitempty"`
	DataType  string `xml:"datatype,attr"`
	Precision int    `xml:"precision,attr,omitempty"`
	Width     int    `xml:"width,attr,omitempty"`
	XType     string `xml:"xtype,attr,omitempty"`
	Ref       string `xml:"ref,attr,omitempty"`
	UCD       string `xml:"ucd,attr,omitempty"`
	UType     string `xml:"utype,attr,omitempty"`
	Value     string `xml:"value,attr"`
	ArraySize string `xml:"arraysize,attr,omitempty"`

	Description string      `xml:"DESCRIPTION,omitempty"`
	Values      []xmlValues `xml:"VALUES,omitempty"`
	Links       []xmlLink   `xml:"LINK,omitempty"`
}

type xmlFieldRef struct {
	Ref string `xml:"ref,attr"`
}

type xmlField struct {
	Name        string   `xml:"name,attr"`
	ID          string   `xml:"ID,attr,omitempty"`
	UCD         string   `xml:"ucd,attr,omitempty"`
	Ref         string   `xml:"ref,attr,omitempty"`
	UType       string   `xml:"utype,attr,omitempty"`
	DataType    string   `xml:"datatype,attr"`
	ArraySize   string   `xml:"arraysize,attr,omitempty"`
	Precision   int      `xml:"precision,attr,omitempty"`
	Width       int      `xml:"width,attr,omitempty"`
	Unit        string   `xml:"unit,attr,omitempty"`
	Description string   `xml:"DESCRIPTION,omitempty"`
	Link        *xmlLink `xml:"LINK,omitempty"`
}

type xmlValues struct {
	Min     *xmlMin     `xml:"MIN,omitempty"`
	Max     *xmlMax     `xml:"MAX,omitempty"`
	Options []xmlOption `xml:"OPTION,omitempty"`

	ID   string `xml:"ID,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"` // default: "legal"
	Null string `xml:"null,attr,omitempty"`
	Ref  string `xml:"ref,attr,omitempty"`
}

type xmlMin struct {
	XMLName   xml.Name `xml:"MIN"`
	Value     string   `xml:"value,attr"`
	Inclusive string   `xml:"inclusive,attr,omitempty"` // default: "yes"
}

type xmlMax struct {
	XMLName   xml.Name `xml:"MAX"`
	Value     string   `xml:"value,attr"`
	Inclusive string   `xml:"inclusive,attr,omitempty"` // default: "yes"
}

type xmlOption struct {
	XMLName xml.Name `xml:"OPTION"`

	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr"`
}

type xmlLink struct {
	ID          string `xml:"ID,attr,omitempty"`
	ContentRole string `xml:"content-role,attr,omitempty"`
	ContentType string `xml:"content-type,attr,omitempty"`
	Title       string `xml:"title,attr,omitempty"`
	Value       string `xml:"value,attr,omitempty"`
	HRef        string `xml:"href,attr,omitempty"`
	GRef        string `xml:"gref,attr,omitempty"`
	Action      string `xml:"action,attr,omitempty"`
}

type xmlRow struct {
	Data [][]byte `xml:"TD"`
}

type xmlTableData struct {
	Rows []xmlRow `xml:"TR"`
}

type xmlBinData struct{}
type xmlBinData2 struct{}
