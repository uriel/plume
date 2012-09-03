package theme

import (
	"html/template"
	"io/ioutil"
)

func Load(name string, path string) (*template.Template, error) {
	b, err := ioutil.ReadFile(path)

	template_str := string(b)
	template_obj, err := template.New(name).Parse(template_str)

	if err != nil {
		return nil, err
	}

	return template_obj, nil
}
