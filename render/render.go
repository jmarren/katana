package render

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
)

func RenderPath(c *src.Component, w http.ResponseWriter, r *http.Request, path string) {
	if r.Header.Get("HX-Request") == "true" {
		currentUrl := r.URL.RawPath
		newUrl, err := ChangePathValue(currentUrl, path, c.Name)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		w.Header().Set("HX-Push-Url", newUrl)
		c.Render(w, r, templates.WrapMerge)
	} else {
		r.SetPathValue(c.Name, path)
		c.Render(w, r, templates.Basefunc)
	}
}

func RenderQuery(c *src.Component, w http.ResponseWriter, r *http.Request, query string) {
	if r.Header.Get("HX-Request") == "true" {
		currentUrl := r.Header.Get("HX-Current-URL")
		newUrl, err := UpdateQueryParam(currentUrl, c.Name, query)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		w.Header().Set("HX-Push-Url", newUrl)
		c.Render(w, r, templates.WrapAppend)
	} else {
		c.Render(w, r, templates.Basefunc)
	}
}

func ChangePathValue(rawUrl string, pathParam string, newValue string) (string, error) {
	fmt.Printf("currentUrl: %s\npathParam: %s\nnewValue: %s\n", rawUrl, pathParam, newValue)
	currUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	fmt.Printf("currUrl: %v\n", currUrl)
	newUrl, err := url.JoinPath(rawUrl, newValue)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	return newUrl, nil
}

func UpdateQueryParam(currentUrl string, key string, value string) (string, error) {
	url, err := url.Parse(currentUrl)
	if err != nil {
		return "", err
	}
	q := url.Query()
	q.Set(key, value)
	url.RawQuery = q.Encode()
	return url.String(), nil
}

// func RenderBase(c *src.Component,w http.ResponseWriter, r *http.Request, baseFunc ParentTemplate) {

// func Render(w http.ResponseWriter, r *http.Request, c *src.Component, bas) {
// 	html := templates.Base(c.Head, c.Body)
// 	html.Render(r.Context(), w)
// }
