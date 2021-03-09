package link

import (
	"io"
)

// Link represent a link (<a href=""></a>) in an html
type Link struct {
	Href string
	Test string
}

// Parse will take html document d will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	// doc, err := html.Parse(r)
	html.
	if err != nil {
		return nil, err
	}
	return nil, nil
}
