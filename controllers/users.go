package controllers

import (
	"net/http"
	"appengine"
	"fmt"
)

type UsersController struct{
	Controller
}

func NewUsersController() *UsersController{
	return &UsersController{
	}
}

func (controller *UsersController) Init(r *http.Request, w http.ResponseWriter){
	controller.R = r
	controller.W = w
	controller.Data = make(map[interface{}]interface{})
	controller.C = appengine.NewContext(r)
	controller.Layout = make([]string,0)
	controller.TemplateNames = ""
	controller.Name = "users"
	controller.DetermineRoutes()
}

func (controller *UsersController) Serve(action string){
	controller.TemplateNames = "Templates"
	if action == "index"{
		controller.Index()
	}else if action == "view"{
		controller.View()
	}else if action =="edit"{
		controller.Edit()
	}else if action =="delete"{
		controller.Delete()
	}
}

func (controller *UsersController) MobileServe(action string){
	
}

func (controller *UsersController) View(){
	controller.Controller.Render(true)
}

func (controller *UsersController) Index(){
	controller.Render(true)
}

func (controller *UsersController) Edit(){
	controller.Render(true)
}

func (controller *UsersController) Delete(){
	controller.Render(true)
}

func (controller *UsersController) Render(show bool) error{
	//controller.Controller.Render(show)
	controller.Controller.Render(show)
	println("Rendering in user")
			fmt.Fprint(controller.Controller.W,"You're over here")
	return nil
}
