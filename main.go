package main

import (
	"ginchat/myrouter"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := myrouter.Router()
	r.Run(":8081")
}
