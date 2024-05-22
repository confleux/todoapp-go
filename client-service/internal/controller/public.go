package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	t, _ := template.ParseGlob("public/template/*.html")

	end := time.Now()

	elapsed := end.Sub(start).Milliseconds()

	data := struct {
		ServerLoadTime int64
	}{
		ServerLoadTime: elapsed,
	}

	err := t.ExecuteTemplate(w, "index", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	t, _ := template.ParseGlob("public/template/*.html")

	end := time.Now()

	elapsed := end.Sub(start).Milliseconds()

	data := struct {
		ServerLoadTime int64
	}{
		ServerLoadTime: elapsed,
	}

	err := t.ExecuteTemplate(w, "form", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	t, _ := template.ParseGlob("public/template/*.html")

	end := time.Now()

	elapsed := end.Sub(start).Milliseconds()

	data := struct {
		ServerLoadTime int64
	}{
		ServerLoadTime: elapsed,
	}

	err := t.ExecuteTemplate(w, "todo", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	t, _ := template.ParseGlob("public/template/*.html")

	end := time.Now()

	elapsed := end.Sub(start).Milliseconds()

	data := struct {
		ServerLoadTime int64
	}{
		ServerLoadTime: elapsed,
	}

	err := t.ExecuteTemplate(w, "login", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	t, _ := template.ParseGlob("public/template/*.html")

	end := time.Now()

	elapsed := end.Sub(start).Milliseconds()

	data := struct {
		ServerLoadTime int64
	}{
		ServerLoadTime: elapsed,
	}

	err := t.ExecuteTemplate(w, "signup", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func TodoAppHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	t, _ := template.ParseGlob("public/template/*.html")

	end := time.Now()

	elapsed := end.Sub(start).Milliseconds()

	data := struct {
		ServerLoadTime int64
	}{
		ServerLoadTime: elapsed,
	}

	err := t.ExecuteTemplate(w, "todo-app", data)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
