package actions

import (
	"strconv"
	"time"

	"github.com/gautamrege/zephyr"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	go zephyr.Setup()
	return c.Render(200, r.HTML("startup.html"))
}

func ComponentHandler(c buffalo.Context) error {
	if c.Param("monitor_rpm") == "true" {
		c.Set("monitor_rpm", true)
	} else {
		c.Set("monitor_rpm", false)
	}

	return c.Render(200, r.HTML("components.html"))
}

func StopHandler(c buffalo.Context) error {
	devices := zephyr.GetDevices()
	devices.PlatterMotor.StopMotor()
	return c.Render(200, r.HTML("components.html"))
}

func MotorHandler(c buffalo.Context) error {
	devices := zephyr.GetDevices()
	speed, _ := strconv.Atoi(c.Param("rpm"))
	devices.PlatterMotor.ChangeSpeed(speed)
	return c.Redirect(302, "/components")
}

func SetupHandler(c buffalo.Context) error {
	zephyr.Setup()
	return c.Redirect(302, "/components")
}

func TracksHandler(c buffalo.Context) error {
	devices := zephyr.GetDevices()

	devices.StartPlatterMotor()
	go devices.DetectTracks()
	return c.Redirect(302, "/components?monitor_rpm=true")
}

func RpmHandler(c buffalo.Context) error {
	wcon, _ := c.Websocket()
	d := zephyr.GetDevices()

	step_time := time.Now()
	for {
		val := <-d.PlatterSpeed.DataChannel
		if val == 0 {
			diff := time.Since(step_time)
			diff_in_ms := int(diff / time.Millisecond)

			if diff_in_ms < 100 {
				continue
			}
			rpm_per_rotation := (1 / diff.Seconds()) * 60
			step_time = time.Now()
			wcon.WriteJSON(map[string]float64{"rpm": rpm_per_rotation})
		}
	}
	return nil
}

func PlayTrackHandler(c buffalo.Context) error {
	devices := zephyr.GetDevices()

	devices.StartPlatterMotor()
	return c.Redirect(302, "/components?monitor_rpm=true")
}
