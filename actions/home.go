package actions

import (
	"github.com/gautamrege/zephyr"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	go zephyr.Setup()
	return c.Render(200, r.HTML("startup.html"))
}
