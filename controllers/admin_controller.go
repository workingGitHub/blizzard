package controllers

import (
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/mvc"
)

// BasicAuth middleware sample.
var AdminBasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "admin",
	},
})

// UsersController is our /users API controller.
// GET				/users  | get all
// GET				/users/{id:long} | get by id
// PUT				/users/{id:long} | update by id
// DELETE			/users/{id:long} | delete by id
// Requires basic authentication.
type AdminController struct {
	mvc.C

	Service services.UserServiceIf
}

// GetLogin handles GET: http://localhost:8080/admin/addcard
func (c *UserController) GetAddcard() mvc.Result {

	var loginStaticView = mvc.View{
		Name: "admin/addcard.html",
		Data: context.Map{"Title": "Add Card"},
	}

	return loginStaticView
}

func (c *UserController) PostAddcard() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
		c.logout()
	}

	var loginStaticView = mvc.View{
		Name: "user/login.html",
		Data: context.Map{"Title": "User Login"},
	}

	return loginStaticView
}
