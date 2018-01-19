package actions

import (
	"strconv"
	"time"

	"github.com/gautamrege/zephyr"
	"github.com/gobuffalo/buffalo"
)

// SettingsReset default implementation.
func SettingsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("settings/index.html"))
}

func ComponentHandler(c buffalo.Context) error {
	if c.Param("monitor_rpm") == "true" {
		c.Set("monitor_rpm", true)
	} else {
		c.Set("monitor_rpm", false)
	}

	return c.Render(200, r.HTML("settings/index.html"))
}

func StopHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	z.Devices.PlatterMotor.StopMotor()
	return c.Render(200, r.HTML("settings/index.html"))
}

func MotorHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	speed, _ := strconv.Atoi(c.Param("rpm"))
	z.Devices.PlatterMotor.ChangePlatterSpeed(speed)
	return c.Redirect(302, "/settings")
}

func SetupHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	z.Restart()
	c.Flash().Add("success", "Initialized the Devices")
	return c.Redirect(302, "/settings")
}

func TracksHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	z.LoadTracks()
	//z.Devices.StartPlatterMotor()
	c.Set("tracks", z.Tracks)
	return c.Render(200, r.HTML("tracks.html"))
}

func RpmHandler(c buffalo.Context) error {
	wcon, _ := c.Websocket()
	z := zephyr.GetInstance()

	step_time := time.Now()
	for {
		val := <-z.Devices.PlatterSpeed.DataChannel
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

func MotorStepHandler(c buffalo.Context) error {
	wcon, _ := c.Websocket()
	z := zephyr.GetInstance()

	for {
		wcon.WriteJSON(map[string]int{"steps": z.Devices.ArmMotor.CurrentPosition})
	}
	return nil
}

func PlayTrackHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	action, _ := strconv.Atoi(c.Param("action"))
	trackNo, _ := strconv.Atoi(c.Param("track"))

	if action == 0 {
		go z.PlayTrack(trackNo)
	} else {
		z.Devices.MoveToneArmForTrack(trackNo)
	}

	return c.Render(200, r.String(""))
}

func ArmSpeedHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	speed, _ := strconv.Atoi(c.Param("speed"))
	z.Devices.ArmMotor.SetArmMotorSpeed(speed)
	return c.Redirect(302, "/settings")
}

func ArmMotorHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	action, _ := strconv.Atoi(c.Param("action"))
	if action == 0 {
		go z.DetectTracks()
	} else {
		z.Devices.ArmMotor.StopMotor()
	}
	return c.Redirect(302, "/settings")
}

func ArmLoaderHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	action, _ := strconv.Atoi(c.Param("action"))
	if action == 0 {
		z.Devices.ArmLoad()
	} else {
		z.Devices.ArmUnload()
	}
	return c.Redirect(302, "/settings")
}

func ArmMotorPWHandler(c buffalo.Context) error {
	z := zephyr.GetInstance()
	pw, _ := strconv.Atoi(c.Request().Form["pw"][0])
	z.Devices.ArmMotor.StepChannel <- pw
	return c.Render(200, r.String(""))
}
