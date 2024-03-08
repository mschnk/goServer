package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mschnk/goServer/routes"
)

func main() {
	routes.CarregaRotas()
http.ListenAndServe(":8080", nil)
}

