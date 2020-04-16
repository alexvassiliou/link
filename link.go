package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link stores the href and inner text values of a HTML document
type Link struct {
	Href string
	Text string
}

// ParseLinks will extract the links from given html
func ParseLinks(r io.Reader) ([]Link, error) {

	z := html.NewTokenizer(r)
	var result []Link
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return result, nil

		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				g := z.Next()
				var l Link
				l.getText(g, z)
				l.getHref(t)
				result = append(result, l)
			}
		}
	}
	return result, nil
}

func (l *Link) getHref(t html.Token) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			l.Href = a.Val
		}
	}
}

func (l *Link) getText(tt html.TokenType, z *html.Tokenizer) {
	if tt == html.TextToken {
		g := z.Token()
		fmt.Println(g)
	}
}
