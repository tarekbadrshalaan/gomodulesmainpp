package main

import (
	"fmt"

	"github.com/tarekbadrshalaan/gomodulesdepone"
)

func main() {
	fmt.Println("Hi from Main app")

	fmt.Println(gomodulesdepone.GetData())

	//using new feature in gomodulesdepone
	fmt.Println(gomodulesdepone.GetExtraData())
}
