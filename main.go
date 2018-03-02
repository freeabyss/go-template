package main

import (
	"maizuo.com/back-end/go-template/src/server/initialize"
)

func main()  {

	initialize.SetupConfig()

	initialize.SetErrorDeal()

	initialize.SetupLogger()

	initialize.SetContext()

	initialize.SetupRedis()

	initialize.SetupDB()

	initialize.SetupNsqProducer()

	//initialize.SetupRPC()

	//timer.SetupTimer()

	initialize.SetupServer()



}
