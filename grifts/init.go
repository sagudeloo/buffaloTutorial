package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sagudeloo/tutorial/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
