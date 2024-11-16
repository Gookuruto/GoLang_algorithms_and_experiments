package main

import (
	"fmt"
	"net/http"
	"workspace/routers"
)

func main() {
	routers.SetupRoutes()

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
