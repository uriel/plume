package theme

import (
	"html/template"
	"io/ioutil"
)

func Load(name string, path string) (tpl *template.Template, err error) {
	b, err := ioutil.ReadFile(path)

	tpl, err = template.New(name).Parse(string(b))

	return
}
