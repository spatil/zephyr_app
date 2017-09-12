package actions

import (
	"github.com/gautamrege/zephyr"
	"github.com/gobuffalo/buffalo"
	"gobot.io/x/gobot/platforms/raspi"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("motor.html"))
}

func MotorHandler(c buffalo.Context) error {
	r := raspi.NewAdaptor()
	zephyr.Monitor_rpm(r, 45.0)
	return c.Redirect(302, "/")
}
