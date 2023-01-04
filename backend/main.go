package main

import "github.com/khalil9022/OJT_Postgre/api"

func main() {
	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
