package index

import (
	"bytes"
	"net/http"
	"strings"
	"github.com/parkn09/quicklish/controllers"
)


func handler(w http.ResponseWriter, r *http.Request){

	params := string(r.URL.Path[1:])
	if(params==""){
		http.ServeFile(w, r, "views/layouts/home.html")
	}else{
	arr := strings.Split(params,"/")
	m := make(map[string]controllers.ControllerInterface)
	m["users"] = controllers.NewUsersController()
	count := len(arr)
	if count > 0{
		if arr[0]=="images"{
			buffer := bytes.NewBufferString("views/")
			buffer.WriteString(string(params))
			http.ServeFile(w, r, buffer.String())
		}else{
			m[arr[0]].Init(r,w)
			action := "index"
			if count >1{
				if arr[1] == "mobile"{
					m[arr[0]].MobileServe(action)
				}else{
					action = arr[1]
					m[arr[0]].Serve(action)
				}
			}else{
				m[arr[0]].Serve(action)
			}
		}
	}	
	}
}


func init(){
	
	//http.HandleFunc("/images*",imageHandler)
	http.HandleFunc("/",handler)

}