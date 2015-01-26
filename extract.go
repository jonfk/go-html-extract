package htmlextractor

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
)

func extractString(data string) (string, err) {

}

func extractUrl(url string) (string, err) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return extractString(body)
}
