package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/spatil/zephyr_app/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
