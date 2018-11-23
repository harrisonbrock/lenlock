package views

import (
	"html/template"
	"net/http"
)

func NewView(layout string, files ...string) *View {

	files = append(files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)

	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// Render the view with predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {

	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
