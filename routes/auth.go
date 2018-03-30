package routes

import (
	"../session"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func GetLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, r *http.Request, s *session.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf(username)
	fmt.Printf(password)

	s.Username = username
	s.IsAuthorized = true

	rnd.Redirect("/")
}

func LogoutHandler(rnd render.Render, r *http.Request, s *session.Session) {

	s.Username = ""
	s.IsAuthorized = false

	rnd.Redirect("/")
}
