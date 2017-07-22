package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("projects/index.html"))
}

// AboutHandler is a default handler to serve the about page.
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}

// LoginHandler is the handler to serve up the login page
func LoginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("login.html"))
}
