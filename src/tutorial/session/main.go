package main

import (
	"html/template"
	"net/http"

	"fmt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9099", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
	}

	var u user
	if uid, ok := dbSessions[c.Value]; ok {
		u = dbUsers[uid]
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("first")
		l := r.FormValue("last")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	fmt.Println(u.UserName)
	tpl.ExecuteTemplate(w, "test.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	e, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	u := dbUsers[e]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
