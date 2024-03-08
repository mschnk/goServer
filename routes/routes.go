package routes

import ("net/http"
"github.com/mschnk/goServer/controllers")


func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
