package routes

import ("net/http"
"github.com/mschnk/goServer/controllers")


func CarregaRotas() {
	http.HandleFunc("/", index)

}
