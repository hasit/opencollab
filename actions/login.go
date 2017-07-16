package actions

import "github.com/gobuffalo/buffalo"

// LoginHandler is the handler to serve up the login page
func LoginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("login.html"))
}
