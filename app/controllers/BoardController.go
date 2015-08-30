package controllers

import "github.com/revel/revel"

type Board struct {
	*revel.Controller
}

func (c Board) Index() revel.Result {
	return c.Render()
}

/*글읽기*/
func (c Board) Read(Id int64) revel.Result {
	return c.Render(Id)
}