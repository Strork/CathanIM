package main

import (
	"furrychat/myrouter"
	"furrychat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := myrouter.Router()
	r.Run(":8081")
}
