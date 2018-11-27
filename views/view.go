package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {

	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)

	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// returns a slice of strings of view file names
func layoutFiles() []string {

	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)

	if err != nil {
		panic(err)
	}

	return files

}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w{} http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// Render the view with predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
