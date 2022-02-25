package main

import (
	"github.com/ianchen0119/go-blogger/web"
)

func main() {
	router := web.NewRouter()
	router.Run(":3000")
}
