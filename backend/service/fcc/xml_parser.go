package fcc

import (
	"bytes"
	"encoding/xml"
	"time"

	"github.com/juju/errors"
	"golang.org/x/net/html/charset"
)

// Filer499QueryResults provides filer query interface
type Filer499QueryResults interface {
	GetFRN() (string, error)
}

// filer499QueryResults implements Filer499QueryResults interface
type filer499QueryResults struct {
	XMLName xml.Name `xml:"Filer499QueryResults"`
	Filers  []filer  `xml:"Filer"`
}

// filer struct
type filer struct {
	XMLName     xml.Name    `xml:"Filer"`
	Form499ID   string      `xml:"Form_499_ID"`
	FilerIDInfo filerIDInfo `xml:"Filer_ID_Info"`
}

// filerIDInfo struct
type filerIDInfo struct {
	XMLName                     xml.Name `xml:"Filer_ID_Info"`
	RegistrationCurrentAsOf     string   `xml:"Registration_Current_as_of"`
	StartDate                   fccDate  `xml:"start_date"`
	USFContributor              string   `xml:"USF_Contributor"`
	LegalName                   string   `xml:"Legal_Name"`
	PrincipalCommunicationsType string   `xml:"Principal_Communications_Type"`
	HoldingCompany              string   `xml:"holding_company"`
	FRN                         string   `xml:"FRN"`
}

type filer499QueryResultsFromXML struct {
	xmlStr string
}

// fccDate struct is custom implementation of date to be able to parse yyyy-mm-dd formal
// FCC API returns dates in yyyy-mm-dd format that default golang XML decoder does not recognize
type fccDate struct {
	time.Time
}

func (c *fccDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02"
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = fccDate{parse}
	return nil
}

// NewFiler499QueryResultsFromXML create new filer query results implementation
func NewFiler499QueryResultsFromXML(xmlStr string) Filer499QueryResults {
	return filer499QueryResultsFromXML{
		xmlStr: xmlStr,
	}
}

func (fq filer499QueryResultsFromXML) GetFRN() (string, error) {
	var fQueryResults filer499QueryResults
	b := []byte(fq.xmlStr)
	reader := bytes.NewReader(b)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&fQueryResults)
	if err != nil {
		return "", errors.Annotatef(err, "failed to unmarshal xml FRN from XML")
	}

	if fQueryResults.Filers == nil || len(fQueryResults.Filers) == 0 {
		return "", errors.New("failed to parse FRN from XML")
	}
	filersResult := fQueryResults.Filers[0]
	filerIDInfo := filersResult.FilerIDInfo
	return filerIDInfo.FRN, nil
}
