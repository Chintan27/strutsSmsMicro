package main

import (
	"fmt"
	"net/http"

	"github.com/Chintan27/strutsSmsMicro/controllers"
	"github.com/Chintan27/strutsSmsMicro/sqldb"
)

func main() {
	db := sqldb.ConnectDB()

	h := controllers.NewBaseHandler(db)

	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()

}
