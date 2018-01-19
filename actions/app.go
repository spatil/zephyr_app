package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_zephyr_app_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Automatically save the session if the underlying
		// Handler does not return an error.
		app.Use(middleware.SessionSaver)

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		if ENV != "test" {
			// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
			// Remove to disable this.
			app.Use(csrf.Middleware)
		}

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/", HomeHandler)
		app.GET("/settings/motor/{rpm}", MotorHandler)
		app.GET("/setup", SetupHandler)
		app.GET("/tracks", TracksHandler)
		app.GET("/settings/stop", StopHandler)
		app.GET("/settings/rpm_monitor", RpmHandler)
		app.GET("/settings/step_monitor", MotorStepHandler)
		app.GET("/settings/play", PlayTrackHandler)
		app.GET("/settings/set_arm_speed/{speed}", ArmSpeedHandler)
		app.GET("/settings/arm_motor/{action}", ArmMotorHandler)
		app.POST("/settings/change_pw", ArmMotorPWHandler)
		app.GET("/settings/loader/{action}", ArmLoaderHandler)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
		app.GET("/settings", SettingsHandler)
		app.GET("/dashboard", DashboardIndex)
		app.GET("/manual_play", ManualPlayHandler)
		app.GET("/platter/{rpm}", PlatterHandler)
		app.GET("/move_arm/{direction}", ArmDirectionHandler)
		app.GET("/restart", RestartHandler)
		app.GET("/play", PlayHandler)
	}

	return app
}
