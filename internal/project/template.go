package project

import (
	"io"
	"text/template"

	"github.com/nicksnyder/xb/internal/fs"
)

type Template struct {
	Template *template.Template
	Data     interface{}
}

func newTemplate(source string, data interface{}) *Template {
	tmpl := template.Must(template.New("").Parse(source))
	return &Template{Template: tmpl, Data: data}
}

func (t *Template) WriteAllTo(w io.Writer) error {
	return t.Template.Execute(w, t.Data)
}

var _ fs.WriterTo = (*Template)(nil)
