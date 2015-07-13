package controllers

import (
	"fmt"

	"iidaPro/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	rows, _ := DbMap.Select(models.User{}, "select * from user")
	for _, row := range rows {
		user := row.(*models.User)
		fmt.Printf("%d, %s, %s\n", user.Id, user.Name, user.HashedPassword)
	}

	return c.Render()
}
