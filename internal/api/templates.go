package api

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	files, err := getHtmlFiles("internal/web/templates")
	if err != nil {
		log.Fatal(err)
	}

	return &Templates{
		templates: template.Must(template.ParseFiles(files...)),
	}
}

func getHtmlFiles(pattern string) ([]string, error) {
	var files []string

	err := filepath.Walk(pattern, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}

		files = append(files, path)

		return err
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
