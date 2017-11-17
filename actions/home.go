package actions

import (
	"strconv"

	"github.com/gautamrege/zephyr"
	"github.com/gobuffalo/buffalo"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("", "startup.html"))
}

func ComponentHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("components.html"))
}

func MotorHandler(c buffalo.Context) error {
	rpm, _ := strconv.ParseFloat(c.Param("rpm"), 64)
	zephyr.TestMotor(rpm, 200)
	return c.Redirect(302, "/components")
}

func SetupHandler(c buffalo.Context) error {
	r := raspi.NewAdaptor()

	robot := gobot.NewRobot("ZephyrBot",
		[]gobot.Connection{r},
	)

	zephyr.Setup(robot)

	go robot.Start()
	return c.Redirect(302, "/components")
}

func TracksHandler(c buffalo.Context) error {
	tracks := zephyr.DetectTracks()
	c.Set("tracks", tracks)
	return c.Render(200, r.HTML("tracks.html"))
}
