package main

import "verifyLinux/routers"

func main() {
	r := routers.InitRouters()
	r.Run(":80")
}
