package actions

import (
	"strconv"

	"github.com/gautamrege/zephyr"
	"github.com/gobuffalo/buffalo"
)

// DashboardIndex default implementation.
func DashboardIndex(c buffalo.Context) error {
	go zephyr.Setup()
	return c.Render(200, r.HTML("dashboard/index.html"))
}

func PlatterHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	speed, _ := strconv.Atoi(c.Param("rpm"))
	z.Devices.PlatterMotor.ChangePlatterSpeed(speed)
	z.Devices.StartPlatterMotor()
	return c.Redirect(302, "/manual_play")
}

func ManualPlayHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("dashboard/manual_play.html"))
}

func ArmDirectionHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	z.Devices.ArmMotor.StopMotor()
	dir, _ := strconv.Atoi(c.Param("direction"))
	z.Devices.MoveToneArm(dir)
	return c.Render(200, r.String(""))
}

func ArmMotorStopHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	z.Devices.ArmMotor.StopMotor()
	return c.Render(200, r.String(""))
}

func PlayHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	go z.PlayTrack(-1)
	return c.Redirect(302, "/manual_play")
}

func RestartHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	z.Restart()
	return c.Redirect(302, "/manual_play")
}
