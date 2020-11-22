package main

import (
	"KubeClipper/routers"
	"fmt"

)

func main() {

	fmt.Println(123)

	router := routers.SetUpRouter()

	_ = router.Run()

	
}
