package fcc

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/juju/errors"
)

const (
	emailHeader = "Contact Email:"
)

// SearchDetailResults provides search detail interface
type SearchDetailResults interface {
	GetEmail() (string, error)
}

// NewSearchDetailFromHTML create new search details implementation
func NewSearchDetailFromHTML(htmlStr string) SearchDetailResults {
	return searchDetailResults{
		htmlStr: htmlStr,
	}
}

// searchDetailResults implements SearchDetailResults interface
type searchDetailResults struct {
	htmlStr string
}

func (sd searchDetailResults) GetEmail() (string, error) {
	var heading string
	var email string

	b := []byte(sd.htmlStr)
	reader := bytes.NewReader(b)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", errors.New("failed to parse email from XML")
	}

	// Find each table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				heading = tableheading.Text()
			})

			if heading == emailHeader {
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					email = tablecell.Text()
				})
			}
		})
	})
	return strings.TrimSpace(email), nil
}
