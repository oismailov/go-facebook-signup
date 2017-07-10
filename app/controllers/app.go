package controllers

import (
	"github.com/revel/revel"
)

type Human struct {
	*revel.Controller
}

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func (c Human) Index() revel.Result {
	users := []User{}
	user := User{"John", 20}
	users = append(users, user)
	users = append(users, User{"Mike", 30})

	return c.RenderJSON(users)
}
