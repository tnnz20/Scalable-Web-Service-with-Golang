package main

import (
	"fmt"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/helpers"
)

func main() {
	fmt.Println("Hello world")

	helpers.RunCronJob()

	// fmt.Println(helpers.RandomNum(1, 100))

}
