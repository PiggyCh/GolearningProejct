package main

import (
	"ginchat/router"
	"ginchat/utils"
)

// gin-swagger middleware
// swagger embed files

// gin-swagger middleware
// swagger embed files
func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8080")
}
