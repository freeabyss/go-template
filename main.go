package main

<<<<<<< HEAD
import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
=======
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
>>>>>>> e62ffc6cb726ed9f8addf64a8fd1efd1297dedc2
