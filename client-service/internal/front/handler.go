package front

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("public/template/*.html")

	err := t.ExecuteTemplate(w, "index", nil)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("public/template/*.html")

	err := t.ExecuteTemplate(w, "form", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("public/template/*.html")

	err := t.ExecuteTemplate(w, "todo", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
