package main

import (
	"linkAja/configs"
	"linkAja/master"
	"log"
)


func main()  {


	db,err,hostServer,portServer:=configs.InitDB()	

	if err != nil{
		log.Fatal(err)
	}
	router := configs.CreatRouter()

	master.InitData(router,db)

	configs.RunServer(router,hostServer,portServer)
}
