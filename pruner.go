package goisgd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


func Shorten(uri string) (string, error) {
	u := "http://is.gd/api.php?longurl=" + url.QueryEscape(uri)

	response, err := http.Get(u)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	shortUri := string(body)

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf(shortUri)
	}
	return shortUri, nil
}
