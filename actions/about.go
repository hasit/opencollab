package actions

import "github.com/gobuffalo/buffalo"

// AboutHandler is a default handler to serve the about page.
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}
