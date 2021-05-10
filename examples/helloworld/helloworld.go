package main

import (
	"log"

	"github.com/geektime007/king"
	"github.com/geektime007/king/transport/http"
)

var (
	Name    = "helloworld"
	Version = "v0.0.1"
)

func main() {

	httpSrv := http.NewServer(http.Address(":8000"))
	//httpSrv.HandlePrefix("/", ),
	//)

	app := king.New(
		king.Name(Name),
		king.Version(Version),
		king.Server(
		//httpSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
