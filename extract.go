package htmlextr

import (
	// "fmt"
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func extractString(data string) (string, error) {
	doc, err := html.Parse(bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	var f func(*html.Node, string)
	f = func(n *html.Node, outer string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			log.Printf("a tag attrs : %#v\n", n.Attr)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, n.Data)
			}
		} else if n.Type == html.ElementNode && n.Data == "p" {
			log.Printf("p tag attrs : %#v\n", n.Attr)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c, n.Data)
			}
		} else {
			log.Printf("other tag : %s tag attrs : %#v\n", n.Data, n.Attr)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, outer)
		}
	}
	f(doc, doc.Data)
	log.Printf("Processing finished")
	// testing
	file, err := os.Create("PageStruct.txt")
	if err != nil {
		log.Fatal(err)
	}
	spew.Fdump(file, doc)

	return "", nil
}

func extractUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return extractString(string(body))
}
