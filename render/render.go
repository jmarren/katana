package render

import (
	"net/http"

	"github.com/jmarren/katana/src"
	"github.com/jmarren/katana/templates"
)

func Render(w http.ResponseWriter, r *http.Request, c *src.Component, bas) {
	html := templates.Base(c.Head, c.Body)
	html.Render(r.Context(), w)
}
