package controllers

import (
	"net/http"
	"appengine"
	"html/template"
	//"path"
)

type Controller struct{
	Name string
	R *http.Request
	W http.ResponseWriter
	Data map[interface{}]interface{}
	Tpl *template.Template
	C appengine.Context
	Layout []string
	TemplateNames string
}

type ControllerInterface interface{
	Init(r *http.Request, w http.ResponseWriter)
	Serve(action string)
	MobileServe(action string)
	Render(show bool) error
}

func NewController() *Controller{
	return &Controller{}
}

func (controller *Controller) Init(r *http.Request, w http.ResponseWriter){
	controller.R = r
	controller.W = w
	controller.Data = make(map[interface{}]interface{})
	controller.C = appengine.NewContext(r)
	controller.Layout = make([]string,0)
	controller.TemplateNames = ""
}

func (controller *Controller) DetermineRoutes(){

	println(controller.R.URL.Path[1:])
}

func (controller *Controller) CheckSession() bool{
	println("Checking session")
	return true
}

func (controller *Controller) Render(show bool) error{
	/*
	if show{
		var filenames []string
        for _, file := range controller.Layout {
            filenames = append(filenames, path.Join("views/", file))
        }
		t,err := template.ParseFiles(filenames...)
		if err != nil{
			return err
		}
		err = t.ExecuteTemplate(controller.W, controller.TemplateNames,controller.Data)
		if err != nil{
			return err
		}
	}
	*/
	println(controller.TemplateNames)
	println("Or here")
	return nil
}

