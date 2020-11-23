package main

import (
	"fmt"
	"github.com/willie-lin/KubeClipper/routers"
)

func main() {

	fmt.Println(123)

	router := routers.SetUpRouter()

	_ = router.Run()

}
