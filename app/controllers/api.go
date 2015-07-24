package controllers

import "github.com/revel/revel"

type Api struct {
	*revel.Controller
}

func (c Api) Index() revel.Result {
	return c.Render()
}

type Event struct {
	Date   string `json:"date"`
	Event  string `json:"event"`
	Num    int    `json:"num"`
	Status int    `json:"status"`
}

func (c Api) Get(num int) revel.Result {
	// fmt.Println(c.Request.PostFormValue("name"))
	if c.Request.Method == "GET" {
		return c.Render()
	}

	date, event := "", ""

	switch num {
	case 0:
		date = "7月25日"
		event = "中間発表"
	case 1:
		date = "7月26日"
		event = "夏休み"
	}
	return c.RenderJson(Event{Date: date, Event: event, Num: num, Status: 200})
}
