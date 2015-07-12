package controllers

import "github.com/revel/revel"

type Api struct {
	*revel.Controller
}

func (c Api) Index() revel.Result {
	return c.Render()
}

type User struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (c Api) Get(name string, status int) revel.Result {
	// fmt.Println(c.Request.PostFormValue("name"))
	if c.Request.Method == "GET" {
		return c.Render()
	}
	return c.RenderJson(User{Name: name, Status: status})
}
