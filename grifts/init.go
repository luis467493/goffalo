package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/products/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
