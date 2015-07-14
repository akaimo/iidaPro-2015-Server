package controllers

import (
	"fmt"
	"iidaPro/app/models"
	"iidaPro/app/routes"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	rows, _ := Dbm.Select(models.User{}, "select * from user")
	for _, row := range rows {
		user := row.(*models.User)
		fmt.Printf("%d, %s, %s\n", user.Id, user.Name, user.HashedPassword)
	}

	return c.Render()
}

func (c App) Login(username, password string) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Events.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.App.Index())
}

func (c App) getUser(username string) *models.User {
	users, err := c.Txn.Select(models.User{}, `select * from user where Name = ?`, username)
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0].(*models.User)
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}
