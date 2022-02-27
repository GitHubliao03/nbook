package main

import (
	"nbook/router"
	_ "nbook/utils"
)

func main() {
	//1.mysql redis
	//2.start router
	router.StartRouter()
	//3.gocorn

}
