package controllers

import (
	"iidaPro/app/routes"

	"github.com/revel/revel"
)

type Events struct {
	App
}

func (c Events) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}
	return nil
}

func (c Events) Index() revel.Result {
	return c.Render()
}
