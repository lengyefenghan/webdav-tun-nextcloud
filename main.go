package main

import (
	"WebdavTunNextcloud/routers"
)

func main() {
	r := routers.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080

}
