package main

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/spatil/zephyr_app/actions"
)

func main() {
	port := envy.Get("PORT", "3000")
	app := actions.App()
	log.Fatal(app.Start(port))
}
