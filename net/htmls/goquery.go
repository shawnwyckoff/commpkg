package htmls

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type (
	SelectorBuilder struct{}
	Doc goquery.Document
)

func NewDocFromHtmlSrc(htmlSrc *string) (*goquery.Document, error) {
	sr := strings.NewReader(*htmlSrc)
	doc, err := goquery.NewDocumentFromReader(sr)
	if err != nil {
		return nil, err
	}
	return doc, nil
}