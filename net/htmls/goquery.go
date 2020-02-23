package htmls

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"strings"
)

func NewDocFromHtmlSrc(htmlSrc *string) (*goquery.Document, error) {
	sr := strings.NewReader(*htmlSrc)
	doc, err := goquery.NewDocumentFromReader(sr)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func NewDocFromHtmlFile(filename string) (*goquery.Document, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	str := string(buf)
	return NewDocFromHtmlSrc(&str)
}

func QueryStringElementWithOneOfClasses(element, oneOfClasses string) string {
	return element + "." + oneOfClasses
}

func QueryStringElementWithClass(element, class string) string {
	return element + "." + class
}
