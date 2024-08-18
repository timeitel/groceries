package pages

import (
	"github.com/a-h/templ"
)

func newURL(s string) string {
	return string(templ.URL(s))
}
